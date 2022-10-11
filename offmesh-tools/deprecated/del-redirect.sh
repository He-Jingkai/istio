PROXY_NAME=$1 #fake
POD_IP=$2 #10.32.0.8
PROXY_IP=$3 #10.32.0.11

# iptables -t nat -D PREROUTING -p tcp -s $POD_IP -j DNAT --to-destination $PROXY_IP:15001

iptables -t nat -D OUTPUT -p tcp -d $POD_IP -j INBOUND_OUTPUT-$PROXY_NAME
iptables -t nat -D PREROUTING -p tcp -d $POD_IP -j INBOUND_PREROUTING-$PROXY_NAME

iptables -t nat -D INBOUND_PREROUTING-$PROXY_NAME -p tcp -s $PROXY_IP -j RETURN
iptables -t nat -D INBOUND_PREROUTING-$PROXY_NAME -p tcp -j IPRULE_REDIRECT-$PROXY_NAME
iptables -t nat -X INBOUND_PREROUTING-$PROXY_NAME

iptables -t nat -D INBOUND_OUTPUT-$PROXY_NAME -p tcp -s $PROXY_IP -j RETURN
iptables -t nat -D INBOUND_OUTPUT-$PROXY_NAME -p tcp -j IPRULE_REDIRECT-$PROXY_NAME
iptables -t nat -X INBOUND_OUTPUT-$PROXY_NAME

iptables -t nat -D IPRULE_REDIRECT-$PROXY_NAME -p tcp --dport 15020 -j DNAT --to-destination $PROXY_IP:15020
iptables -t nat -D IPRULE_REDIRECT-$PROXY_NAME -p tcp --dport 15021 -j DNAT --to-destination $PROXY_IP:15021
iptables -t nat -D IPRULE_REDIRECT-$PROXY_NAME -p tcp -j DNAT --to-destination $PROXY_IP:15006
iptables -t nat -X IPRULE_REDIRECT-$PROXY_NAME




