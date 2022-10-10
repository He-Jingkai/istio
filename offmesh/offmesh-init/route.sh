# The purpose of this script is to route
#       all traffic entering this pod (excluding traffic from the proxy)
#   and all traffic from this pod
# to the proxy pod.

POD_IP=$1
PROXY_IP=$2
# istio-init iptables:
#     -P PREROUTING ACCEPT
#     -P INPUT ACCEPT
#     -P OUTPUT ACCEPT
#     -P POSTROUTING ACCEPT
#     -N ISTIO_INBOUND
#     -N ISTIO_IN_REDIRECT
#     -N ISTIO_OUTPUT
#     -N ISTIO_REDIRECT
#     -A PREROUTING -p tcp -j ISTIO_INBOUND
#     -A OUTPUT -p tcp -j ISTIO_OUTPUT
#     -A ISTIO_INBOUND -p tcp -m tcp --dport 22 -j RETURN
#     -A ISTIO_INBOUND -p tcp -m tcp --dport 15020 -j RETURN
#     -A ISTIO_INBOUND -p tcp -j ISTIO_IN_REDIRECT
#     -A ISTIO_IN_REDIRECT -p tcp -j REDIRECT --to-ports 15006
#     -A ISTIO_OUTPUT -s 127.0.0.6/32 -o lo -j RETURN
#     -A ISTIO_OUTPUT ! -d 127.0.0.1/32 -o lo -j ISTIO_IN_REDIRECT
#     -A ISTIO_OUTPUT -m owner --uid-owner 1337 -j RETURN
#     -A ISTIO_OUTPUT -m owner --gid-owner 1337 -j RETURN
#-A ISTIO_OUTPUT -d 127.0.0.1/32 -j RETURN
#-A ISTIO_OUTPUT -j ISTIO_REDIRECT
#-A ISTIO_REDIRECT -p tcp -j REDIRECT --to-ports 15001

# STEP 1: mark all the pockets to route

# outbound traffic
iptables -t mangle -A PREROUTING -p tcp -s $POD_IP -j MARK --set-mark 150

# inbound traffic
iptables -t mangle -N REDIRECT_MARK
iptables -t mangle -A REDIRECT_MARK -p tcp -j MARK --set-mark 150

iptables -t mangle -N INBOUND_PREROUTING
iptables -t mangle -A INBOUND_PREROUTING -p tcp -s $PROXY_IP -j RETURN
iptables -t mangle -A INBOUND_PREROUTING -p tcp -j REDIRECT_MARK

iptables -t mangle -N INBOUND_OUTPUT
iptables -t mangle -A INBOUND_OUTPUT -p tcp -s $PROXY_IP -j RETURN
iptables -t mangle -A INBOUND_OUTPUT -p tcp -j REDIRECT_MARK

iptables -t mangle -A PREROUTING -p tcp -j INBOUND_PREROUTING
iptables -t mangle -A OUTPUT -p tcp -j INBOUND_OUTPUT

