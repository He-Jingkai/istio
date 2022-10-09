PROXY_NAME=$1 #test
POD_IP=$2 #10.32.0.6
PROXY_IP=$3 #10.32.0.8


iptables -t nat -A PREROUTING -p tcp -s $POD_IP -j DNAT --to-destination $PROXY_IP:15001

iptables -t nat -N IPRULE_REDIRECT-$PROXY_NAME
iptables -t nat -A IPRULE_REDIRECT-$PROXY_NAME -p tcp --dport 15020 -j DNAT --to-destination $PROXY_IP:15020
iptables -t nat -A IPRULE_REDIRECT-$PROXY_NAME -p tcp --dport 15021 -j DNAT --to-destination $PROXY_IP:15021
iptables -t nat -A IPRULE_REDIRECT-$PROXY_NAME -p tcp -j DNAT --to-destination $PROXY_IP:15006

iptables -t nat -N INBOUND_PREROUTING-$PROXY_NAME
iptables -t nat -A INBOUND_PREROUTING-$PROXY_NAME -p tcp -s $PROXY_IP -j RETURN
iptables -t nat -A INBOUND_PREROUTING-$PROXY_NAME -p tcp -j IPRULE_REDIRECT-$PROXY_NAME

iptables -t nat -N INBOUND_OUTPUT-$PROXY_NAME
iptables -t nat -A INBOUND_OUTPUT-$PROXY_NAME -p tcp -s $PROXY_IP -j RETURN
iptables -t nat -A INBOUND_OUTPUT-$PROXY_NAME -p tcp -j IPRULE_REDIRECT-$PROXY_NAME

iptables -t nat -A OUTPUT -p tcp -d $POD_IP -j INBOUND_OUTPUT-$PROXY_NAME
iptables -t nat -A PREROUTING -p tcp -d $POD_IP -j INBOUND_PREROUTING-$PROXY_NAME






#但是envoy中的filter针对sidecar的定制化程度过高