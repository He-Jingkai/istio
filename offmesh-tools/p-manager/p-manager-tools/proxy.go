package p_manager_tools

import (
	"context"
	"github.com/google/uuid"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

// all the proxy's will be put on namespace ProxyNamespace
/*Volumes:
 */
const ProxyNamespace = `offmesh-istio-proxy`

type PodMeta struct {
	NameSpace string
	Name      string
}

func CreateNewProxy(clientSet *kubernetes.Clientset) (*PodMeta, error) {
	proxyName := `proxy-` + uuid.New().String()
	cpuLimit, _ := resource.ParseQuantity(`2`)
	memoryLimit, _ := resource.ParseQuantity(`1Gi`)
	cpuRequest, _ := resource.ParseQuantity(`10m`)
	memoryRequest, _ := resource.ParseQuantity(`40Mi`)
	val420 := int32(420) //TODO:两个Volume共用一个是否会有问题
	val43200 := int64(43200)
	val3607 := int64(3607)

	//TODO: 根据istio config文件自动生成
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
			Containers: []corev1.Container{{
				Image: `docker.io/istio/proxyv2:1.15.0`,
				Resources: corev1.ResourceRequirements{
					Limits:   corev1.ResourceList{corev1.ResourceLimitsCPU: cpuLimit, corev1.ResourceLimitsMemory: memoryLimit},
					Requests: corev1.ResourceList{corev1.ResourceRequestsCPU: cpuRequest, corev1.ResourceRequestsMemory: memoryRequest},
				},
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
					`--proxyLogLevel=warning`,
					`--proxyComponentLogLevel=misc:error`,
					`--log_output_level=default:info`,
					`--concurrency`,
					`2`,
				},
				VolumeMounts: []corev1.VolumeMount{{
					Name:      `istio-podinfo`,
					MountPath: `/etc/istio/pod`,
				}, {
					Name:      `istio-envoy`,
					MountPath: `/etc/istio/proxy`,
				}, {
					Name:      `istio-data`,
					MountPath: `/var/lib/istio/data`,
				}, {
					Name:      `credential-socket`,
					MountPath: `/var/run/secrets/credential-uds`,
				}, {
					Name:      `istiod-ca-cert`,
					MountPath: `/var/run/secrets/istio`,
				}, {
					Name:      `kube-api-access-5xbm5`,
					MountPath: `/var/run/secrets/kubernetes.io/serviceaccount`,
				}, {
					Name:      `istio-token`,
					MountPath: `/var/run/secrets/tokens`,
				}, {
					Name:      `workload-certs`,
					MountPath: `/var/run/secrets/workload-spiffe-credentials`,
				}, {
					Name:      `workload-socket`,
					MountPath: `/var/run/secrets/workload-spiffe-uds`,
				},
				},
				//TODO: Environment and Port
				Env:   []corev1.EnvVar{},
				Ports: []corev1.ContainerPort{},
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
				{
					Name: `kube-api-access-5xbm5`,
					VolumeSource: corev1.VolumeSource{
						Projected: &corev1.ProjectedVolumeSource{
							DefaultMode: &val420,
							Sources: []corev1.VolumeProjection{
								{
									ServiceAccountToken: &corev1.ServiceAccountTokenProjection{
										ExpirationSeconds: &val3607,
										Path:              `token`,
									},
								},
								{
									ConfigMap: &corev1.ConfigMapProjection{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: `kube-root-ca.crt`,
										},
										Items: []corev1.KeyToPath{{Key: `ca.crt`, Path: `ca.crt`}},
									},
								},
								{
									DownwardAPI: &corev1.DownwardAPIProjection{
										Items: []corev1.DownwardAPIVolumeFile{
											{FieldRef: &corev1.ObjectFieldSelector{FieldPath: `metadata.namespace`}},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	_, err := clientSet.CoreV1().Pods(ProxyNamespace).Create(context.Background(), newPod, metav1.CreateOptions{})
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
