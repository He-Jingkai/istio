iptables -t mangle -A OUTPUT -p tcp -j MARK --set-mark 100

ip route add default via $PROXY_IP table 100
ip rule add fwmark 100 table 100