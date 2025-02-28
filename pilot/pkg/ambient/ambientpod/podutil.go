// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ambientpod

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"istio.io/api/label"
	"istio.io/api/mesh/v1alpha1"
	"istio.io/istio/pilot/pkg/ambient"
	"istio.io/pkg/log"
)

func WorkloadFromPod(pod *corev1.Pod) ambient.Workload {
	var containers, ips []string
	for _, container := range pod.Spec.Containers {
		containers = append(containers, container.Name)
	}
	for _, ip := range pod.Status.PodIPs {
		ips = append(ips, ip.IP)
	}

	var controllerName, controllerKind string
	for _, ref := range pod.GetOwnerReferences() {
		if ref.Controller != nil && *ref.Controller {
			controllerName, controllerKind = ref.Name, ref.Kind
			break
		}
	}

	return ambient.Workload{
		UID:               string(pod.UID),
		Name:              pod.Name,
		Namespace:         pod.Namespace,
		Labels:            pod.Labels, // TODO copy?
		ServiceAccount:    pod.Spec.ServiceAccountName,
		NodeName:          pod.Spec.NodeName,
		HostNetwork:       pod.Spec.HostNetwork,
		PodIP:             pod.Status.PodIP,
		PodIPs:            ips,
		CreationTimestamp: pod.CreationTimestamp.Time,
		WorkloadMetadata: ambient.WorkloadMetadata{
			GenerateName:   pod.GenerateName,
			Containers:     containers,
			ControllerName: controllerName,
			ControllerKind: controllerKind,
		},
	}
}

func hasPodIP(pod *corev1.Pod) bool {
	return pod.Status.PodIP != ""
}

func isRunning(pod *corev1.Pod) bool {
	return pod.Status.Phase == corev1.PodRunning
}

func ShouldPodBeInIpset(namespace *corev1.Namespace, pod *corev1.Pod, meshMode string, ignoreNotRunning bool) bool {
	// Pod must:
	// - Be running
	// - Have an IP address
	// - Ambient mesh not be off
	// - Cannot have a legacy label (istio.io/rev or istio-injection=enabled)
	// - If mesh is in namespace mode, must be in active namespace
	if (ignoreNotRunning || (isRunning(pod) && hasPodIP(pod))) &&
		meshMode != AmbientMeshOff.String() &&
		!HasLegacyLabel(pod.GetLabels()) &&
		!PodHasOptOut(pod) &&
		IsNamespaceActive(namespace, meshMode) {
		return true
	}

	return false
}

// @TODO Interim function for waypoint proxy, to be replaced after design meeting
func PodHasOptOut(pod *corev1.Pod) bool {
	if val, ok := pod.Labels["ambient-type"]; ok {
		return val == "waypoint" || val == "none"
	}
	return false
}

func IsNamespaceActive(namespace *corev1.Namespace, meshMode string) bool {
	// Must:
	// - MeshConfig be in an "ON" mode
	// - MeshConfig must be in a "DEFAULT" mode, plus:
	//   - Namespace cannot have "legacy" labels (ie. istio.io/rev or istio-injection=enabled)
	//   - Namespace must have label istio.io/dataplane-mode=ambient
	if meshMode == AmbientMeshOn.String() ||
		(meshMode == AmbientMeshNamespace.String() &&
			namespace != nil &&
			!HasLegacyLabel(namespace.GetLabels()) &&
			namespace.GetLabels()["istio.io/dataplane-mode"] == "ambient") {
		return true
	}

	return false
}

func HasSelectors(lbls map[string]string, selectors []*v1.LabelSelector) bool {
	for _, selector := range selectors {
		sel, err := v1.LabelSelectorAsSelector(selector)
		if err != nil {
			log.Errorf("Failed to parse selector: %v", err)
			return false
		}

		if sel.Matches(labels.Set(lbls)) {
			return true
		}
	}
	return false
}

var LegacySelectors = []*v1.LabelSelector{
	{
		MatchExpressions: []v1.LabelSelectorRequirement{
			{
				Key:      "istio-injection",
				Operator: v1.LabelSelectorOpIn,
				Values: []string{
					"enabled",
				},
			},
		},
	},
	{
		MatchExpressions: []v1.LabelSelectorRequirement{
			{
				Key:      label.IoIstioRev.Name,
				Operator: v1.LabelSelectorOpExists,
			},
		},
	},
}

// We do not support the istio.io/rev or istio-injection sidecar labels
// If a pod or namespace has these labels, ambient mesh will not be applied
// to that namespace
func HasLegacyLabel(lbl map[string]string) bool {
	for _, ls := range LegacySelectors {
		sel, err := v1.LabelSelectorAsSelector(ls)
		if err != nil {
			log.Errorf("Failed to parse legacy selector: %v", err)
			return false
		}

		if sel.Matches(labels.Set(lbl)) {
			return true
		}
	}

	return false
}

const (
	AmbientMeshNamespace = v1alpha1.MeshConfig_AmbientMeshConfig_DEFAULT
	AmbientMeshOff       = v1alpha1.MeshConfig_AmbientMeshConfig_OFF
	AmbientMeshOn        = v1alpha1.MeshConfig_AmbientMeshConfig_ON
)
