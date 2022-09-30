kubectl get svc --all-namespaces
kubectl get pods --all-namespaces
curl http://10.32.0.12/distribute_proxy/iprule-test/iprule-client-pod
kubectl apply -f p_manager_service.yaml
kubectl delete -f p_manager_service.yaml
kubectl delete pod  -n offmesh-istio-proxy
kubectl logs -n offmesh-istio-proxy 