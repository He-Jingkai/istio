docker build -t p_manager .
docker run -it -p 80:80 p_manager
docker ps -a
docker commit 9e4bf23708fe hejingkai/p_manager
docker push hejingkai/p_manager
# hejingkai/p_manager

kubectl get svc --all-namespaces
kubectl get pods --all-namespaces
curl http://10.109.82.179/distribute_proxy/iprule-test/iprule-client-pod
kubectl apply -f p-manager-svc.yaml
kubectl delete -f p-manager-svc.yaml
kubectl delete pod -n offmesh-istio-proxy
kubectl logs -n offmesh-istio-proxy 
kubectl label namespace iprule-test istio-injection=disabled --overwrite
kubectl label namespace iprule-test istio-injection=enabled --overwrite