package p_manager_tools

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	corev1 "k8s.io/api/core/v1"
	// "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

// all the proxy's will be put on namespace ProxyNamespace

const ProxyNamespace = `offmesh-istio-proxy`

type PodMeta struct {
	NameSpace string
	Name      string
}

func CreateNewProxy(pod *PodMeta, clientSet *kubernetes.Clientset) (*PodMeta, error) {
	podInfo, err := clientSet.CoreV1().Pods(pod.NameSpace).Get(context.Background(), pod.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	proxyName := `proxy-` + uuid.New().String()
	// cpuLimit, _ := resource.ParseQuantity(`2`)
	// memoryLimit, _ := resource.ParseQuantity(`1Gi`)
	// cpuRequest, _ := resource.ParseQuantity(`10m`)
	// memoryRequest, _ := resource.ParseQuantity(`40Mi`)
	val420 := int32(420) //TODO:两个Volume共用一个是否会有问题
	val43200 := int64(43200)
	val1337 := int64(1337)
	valTure := true
	valFalse := false
	ISTIO_META_APP_CONTAINERS := ``
	for _, container := range podInfo.Spec.Containers {
		if ISTIO_META_APP_CONTAINERS != `` {
			ISTIO_META_APP_CONTAINERS += `,`
		}
		ISTIO_META_APP_CONTAINERS += container.Name
	}

	newPod := &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			APIVersion: `v1`,
			Kind:       `Pod`,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      proxyName,
			Namespace: ProxyNamespace,
		},
		Spec: corev1.PodSpec{
			NodeName: podInfo.Spec.NodeName,
			// InitContainers: []corev1.Container{{
			// 	Image: `docker.io/istio/proxyv2:1.15.0`,
			// 	Name:  `istio-validation`,
			// 	// Resources: corev1.ResourceRequirements{
			// 	// 	Limits:   corev1.ResourceList{corev1.ResourceLimitsCPU: cpuLimit, corev1.ResourceLimitsMemory: memoryLimit},
			// 	// 	Requests: corev1.ResourceList{corev1.ResourceRequestsCPU: cpuRequest, corev1.ResourceRequestsMemory: memoryRequest},
			// 	// },
			// 	SecurityContext: &corev1.SecurityContext{
			// 		Capabilities:             &corev1.Capabilities{Drop: []corev1.Capability{`ALL`}},
			// 		RunAsGroup:               &val1337,
			// 		RunAsUser:                &val1337,
			// 		RunAsNonRoot:             &valTure,
			// 		Privileged:               &valFalse,
			// 		ReadOnlyRootFilesystem:   &valTure,
			// 		AllowPrivilegeEscalation: &valFalse,
			// 	},
			// 	Args: []string{
			// 		`istio-iptables`,
			// 		`-p`,
			// 		`15001`,
			// 		`-z`,
			// 		`15006`,
			// 		`-u`,
			// 		`1337`,
			// 		`-m`,
			// 		`REDIRECT`,
			// 		`-i`,
			// 		`*`,
			// 		`-x`,
			// 		``,
			// 		`-b`,
			// 		`*`,
			// 		`-d`,
			// 		`15090,15021,15020`,
			// 		`--log_output_level=default:info`,
			// 		`--run-validation`,
			// 		`--skip-rule-apply`,
			// 	},
			// }},
			Containers: []corev1.Container{{
				Image: `docker.io/hejingkai/proxyv2:1.15-dev`,
				Name:  `istio-proxy`,
				// Resources: corev1.ResourceRequirements{
				// 	Limits:   corev1.ResourceList{corev1.ResourceLimitsCPU: cpuLimit, corev1.ResourceLimitsMemory: memoryLimit},
				// 	Requests: corev1.ResourceList{corev1.ResourceRequestsCPU: cpuRequest, corev1.ResourceRequestsMemory: memoryRequest},
				// },
				ReadinessProbe: &corev1.Probe{
					ProbeHandler: corev1.ProbeHandler{
						HTTPGet: &corev1.HTTPGetAction{
							Path: `/healthz/ready`,
							Port: intstr.Parse(`15021`),
						}},
					InitialDelaySeconds: 1,
					TimeoutSeconds:      3,
					PeriodSeconds:       2,
					SuccessThreshold:    1,
					FailureThreshold:    30,
				},
				Args: []string{
					`proxy`,
					`sidecar`,
					`--domain`,
					`$(POD_NAMESPACE).svc.cluster.local`,
					`--proxyLogLevel=trace`,
					`--proxyComponentLogLevel=misc:error`,
					`--log_output_level=default:info`,
					`--concurrency`,
					`2`,
				},
				VolumeMounts: []corev1.VolumeMount{
					{
						Name:      `istio-podinfo`,
						MountPath: `/etc/istio/pod`,
					},
					{
						Name:      `istio-envoy`,
						MountPath: `/etc/istio/proxy`,
					},
					{
						Name:      `istio-data`,
						MountPath: `/var/lib/istio/data`,
					},
					{
						Name:      `credential-socket`,
						MountPath: `/var/run/secrets/workload-uds`,
					},
					{
						Name:      `istiod-ca-cert`,
						MountPath: `/var/run/secrets/istio`,
					},
					{
						Name:      `istio-token`,
						MountPath: `/var/run/secrets/tokens`,
					},
					{
						Name:      `workload-certs`,
						MountPath: `/var/run/secrets/workload-spiffe-credentials`,
					},
					{
						Name:      `workload-socket`,
						MountPath: `/var/run/secrets/workload-spiffe-uds`,
					},
				},
				Env: []corev1.EnvVar{
					{Name: `JWT_POLICY`, Value: `third-party-jwt`},
					{Name: `PILOT_CERT_PROVIDER`, Value: `istiod`},
					{Name: `CA_ADDR`, Value: `istiod.istio-system.svc:15012`},
					{Name: `POD_NAME`, Value: pod.Name},
					{Name: `POD_NAMESPACE`, Value: pod.NameSpace},
					{Name: `INSTANCE_IP`, Value: podInfo.Status.PodIP},
					{Name: `SERVICE_ACCOUNT`, Value: `default`},
					{Name: `HOST_IP`, Value: podInfo.Status.HostIP},
					{Name: `PROXY_CONFIG`, Value: `{}`},
					{Name: `ISTIO_META_APP_CONTAINERS`, Value: ISTIO_META_APP_CONTAINERS},
					{Name: `ISTIO_META_CLUSTER_ID`, Value: `Kubernetes`},
					{Name: `ISTIO_META_INTERCEPTION_MODE`, Value: `REDIRECT`},
					{Name: `ISTIO_META_WORKLOAD_NAME`, Value: pod.Name},
					{Name: `ISTIO_META_OWNER`, Value: fmt.Sprintf("kubernetes://apis/v1/namespaces/%s/pods/%s", pod.NameSpace, pod.Name)},
					{Name: `ISTIO_META_MESH_ID`, Value: `cluster.local`},
					{Name: `TRUST_DOMAIN`, Value: `cluster.local`},
				},
				Ports: []corev1.ContainerPort{
					{Name: `http-envoy-prom`, ContainerPort: 15090},
					{Name: `inbound`, ContainerPort: 15006},
					{Name: `outbound`, ContainerPort: 15001},
				},
				SecurityContext: &corev1.SecurityContext{
					Capabilities:             &corev1.Capabilities{Drop: []corev1.Capability{`ALL`}},
					RunAsGroup:               &val1337,
					RunAsUser:                &val1337,
					RunAsNonRoot:             &valTure,
					Privileged:               &valTure,
					ReadOnlyRootFilesystem:   &valTure,
					AllowPrivilegeEscalation: &valFalse,
				},
			}},
			Volumes: []corev1.Volume{
				{
					Name:         `workload-socket`,
					VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}},
				},
				{
					Name:         `credential-socket`,
					VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}},
				},
				{
					Name:         `workload-certs`,
					VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}},
				},
				{
					Name:         `istio-envoy`,
					VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{Medium: `Memory`}},
				},
				{
					Name:         `istio-data`,
					VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}},
				},
				{
					Name: `istio-podinfo`,
					VolumeSource: corev1.VolumeSource{
						DownwardAPI: &corev1.DownwardAPIVolumeSource{
							DefaultMode: &val420,
							Items: []corev1.DownwardAPIVolumeFile{
								{Path: `labels`, FieldRef: &corev1.ObjectFieldSelector{FieldPath: `metadata.labels`}},
								{Path: `annotations`, FieldRef: &corev1.ObjectFieldSelector{FieldPath: `metadata.annotations`}}},
						},
					},
				},
				{
					Name: `istio-token`,
					VolumeSource: corev1.VolumeSource{Projected: &corev1.ProjectedVolumeSource{
						DefaultMode: &val420,
						Sources: []corev1.VolumeProjection{
							{ServiceAccountToken: &corev1.ServiceAccountTokenProjection{
								Audience:          `istio-ca`,
								ExpirationSeconds: &val43200,
								Path:              `istio-token`,
							}},
						},
					}},
				},
				{
					Name: `istiod-ca-cert`,
					VolumeSource: corev1.VolumeSource{
						ConfigMap: &corev1.ConfigMapVolumeSource{
							DefaultMode: &val420,
							LocalObjectReference: corev1.LocalObjectReference{
								Name: `istio-ca-root-cert`,
							},
						},
					},
				},
			},
			RestartPolicy: corev1.RestartPolicyAlways,
		},
	}
	_, err = clientSet.CoreV1().Pods(ProxyNamespace).Create(context.Background(), newPod, metav1.CreateOptions{})
	if err != nil {
		return &PodMeta{}, err
	}
	podMeta := PodMeta{
		NameSpace: ProxyNamespace,
		Name:      proxyName,
	}
	return &podMeta, nil
}

func DeleteProxy(clientSet *kubernetes.Clientset, podName string) error {
	return clientSet.CoreV1().Pods(ProxyNamespace).Delete(context.Background(), podName, metav1.DeleteOptions{})
}
