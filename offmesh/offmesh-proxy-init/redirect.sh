POD_IP=$1
PROXY_IP=`tail -n 1 /etc/hosts | awk '{print $1}'`
echo POD_IP=${POD_IP}
echo PROXY_IP=${PROXY_IP}

# inbound (dst = micro-svc pod's ip)
iptables -t nat -N ISTIO_INBOUND
iptables -t nat -A ISTIO_INBOUND -p tcp --dport 15020 -j REDIRECT --to 15020
iptables -t nat -A ISTIO_INBOUND -p tcp --dport 15021 -j REDIRECT --to 15021
iptables -t nat -A ISTIO_INBOUND -p tcp -j REDIRECT --to 15006

iptables -t nat -A PREROUTING -p tcp -d $POD_IP -j ISTIO_INBOUND

# outbound (src = micro-svc pod's ip)
iptables -t nat -A PREROUTING -s $POD_IP ! -d $PROXY_IP -p tcp -j REDIRECT --to 15001

echo "--- iptables ---"
iptables -t nat -L -n --line-numbers
