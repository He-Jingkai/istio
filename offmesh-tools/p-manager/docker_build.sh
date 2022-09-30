docker build -t p_manager .
docker run -it -p 80:80 p_manager

# hejingkai/p_manager

# 2022-09-30T06:27:25.036020Z     info    xdsproxy        connected to upstream XDS server: istiod.istio-system.svc:15012
# 2022-09-30T06:27:25.038694Z     warn    xdsproxy        upstream [6] terminated with unexpected error rpc error: 
# code = PermissionDenied desc = authorization failed: no identities ([spiffe://cluster.local/ns/offmesh-istio-proxy/sa/default]) 
# matched iprule-test/default
# 2022-09-30T06:27:25.039010Z     warning envoy config    StreamAggregatedResources gRPC config stream closed: 7, 
# authorization failed: no identities ([spiffe://cluster.local/ns/offmesh-istio-proxy/sa/default]) matched iprule-test/default