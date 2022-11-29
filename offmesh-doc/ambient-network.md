# ambient network notes

```shell
kubectl create configmap offmesh-conf -n istio-system --from-file=$(pwd)/offmesh-conf
kubectl create configmap offmesh-conf -n kube-system  --from-file=/home/hjk/offmesh-conf

kubectl label namespace default istio.io/dataplane-mode=ambient
```

```yaml
volumeMounts:
  - mountPath: /etc/offmesh-conf
    name: offmesh-conf
      
volumes:
  - name: offmesh-conf
    configMap:
      name: offmesh-conf
```

```shell
#cpu 
kubectl label nodes "$NodeName" offMeshNodeType=cpu
#dpu
kubectl label nodes "$NodeName" offMeshNodeType=dpu
```