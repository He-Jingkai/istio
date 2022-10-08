kubectl get svc --all-namespaces
kubectl get pods --all-namespaces
curl http://10.110.142.4/distribute_proxy/iprule-test/iprule-client-pod
kubectl apply -f p-manager-svc.yaml
kubectl delete -f p-manager-svc.yaml
kubectl delete pod -n offmesh-istio-proxy
kubectl logs -n offmesh-istio-proxy 
kubectl label namespace iprule-test istio-injection=disabled --overwrite
kubectl label namespace iprule-test istio-injection=enabled --overwrite