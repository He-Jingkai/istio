# Script Explain
The purpose of this script is to redirect
 - all traffic to micro-svc pod routed to this pod
 - all traffic from micro-svc pod routed to this pod
to corresponding port.

# istio-init iptables:
```shell
-P PREROUTING ACCEPT
-P INPUT ACCEPT
-P OUTPUT ACCEPT
-P POSTROUTING ACCEPT

-N ISTIO_INBOUND
-A ISTIO_INBOUND -p tcp -m tcp --dport 22 -j RETURN # currently not support
-A ISTIO_INBOUND -p tcp -m tcp --dport 15020 -j RETURN
-A ISTIO_INBOUND -p tcp -j ISTIO_IN_REDIRECT
-N ISTIO_IN_REDIRECT
-A ISTIO_IN_REDIRECT -p tcp -j REDIRECT --to-ports 15006

-A PREROUTING -p tcp -j ISTIO_INBOUND

-N ISTIO_OUTPUT
-A ISTIO_OUTPUT -s 127.0.0.6/32 -o lo -j RETURN # currently no need
-A ISTIO_OUTPUT -d 127.0.0.1/32 -j RETURN # currently no need
-A ISTIO_OUTPUT ! -d 127.0.0.1/32 -o lo -j ISTIO_IN_REDIRECT # currently no need
-A ISTIO_OUTPUT -m owner --uid-owner 1337 -j RETURN # currently no need
-A ISTIO_OUTPUT -m owner --gid-owner 1337 -j RETURN # currently no need
-A ISTIO_OUTPUT -j ISTIO_REDIRECT
-N ISTIO_REDIRECT
-A ISTIO_REDIRECT -p tcp -j REDIRECT --to-ports 15001
-A OUTPUT -p tcp -j ISTIO_OUTPUT
```
# offmesh-proxy-init iptables:

All the following rules only need to be added to the prerouting chain
```shell
# inbound (dst = micro-svc pod's ip)
iptables -t nat -N ISTIO_INBOUND
iptables -t nat -A ISTIO_INBOUND -p tcp --dport 15020 -j REDIRECT --to 15020
iptables -t nat -A ISTIO_INBOUND -p tcp --dport 15021 -j REDIRECT --to 15021
iptables -t nat -A ISTIO_INBOUND -p tcp -j REDIRECT --to 15006

iptables -t nat -A PREROUTING -p tcp -d $POD_IP -j ISTIO_INBOUND

# outbound (src = micro-svc pod's ip)
iptables -t nat -A PREROUTING -s $POD_IP -p tcp -j REDIRECT --to 15001
```

