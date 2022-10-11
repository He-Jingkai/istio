NAME=$1
POD_IP=$2
PROXY_IP=$3
TABLE_NUM=$4

# outbound mark
iptables -t mangle -D PREROUTING -p tcp -s $POD_IP -j MARK --set-mark $TABLE_NUM
# inbound mark
iptables -t mangle -D OUTPUT -p tcp -d $POD_IP -j IN-$NAME
iptables -t mangle -D PREROUTING -p tcp -d $POD_IP -j IN-$NAME

iptables -t mangle -D IN-$NAME -p tcp -s $PROXY_IP -j RETURN
iptables -t mangle -D IN-$NAME -p tcp -j MARK --set-mark $TABLE_NUM
iptables -t mangle -X IN-$NAME

# add route
ip route del default via $PROXY_IP table $TABLE_NUM
ip rule del fwmark $TABLE_NUM table $TABLE_NUM
