NAME=$1
POD_IP=$2 #10.32.0.5
PROXY_IP=$3 #10.32.0.6
TABLE_NUM=$4

# outbound mark
iptables -t mangle -A PREROUTING -p tcp -s $POD_IP -j MARK --set-mark $TABLE_NUM
# inbound mark
iptables -t mangle -N IN-$NAME
iptables -t mangle -A IN-$NAME -p tcp -s $PROXY_IP -j RETURN
iptables -t mangle -A IN-$NAME -p tcp -j MARK --set-mark $TABLE_NUM

iptables -t mangle -A OUTPUT -p tcp -d $POD_IP -j IN-$NAME
iptables -t mangle -A PREROUTING -p tcp -d $POD_IP -j IN-$NAME
# add route
ip route add default via $PROXY_IP table $TABLE_NUM
ip rule add fwmark $TABLE_NUM table $TABLE_NUM

ip route flush cache

