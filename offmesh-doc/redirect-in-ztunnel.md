```shell
$ iptables-save
*filter
:INPUT ACCEPT [13872:25630905]
:FORWARD ACCEPT [0:0]
:OUTPUT ACCEPT [8337:431154]
-A INPUT -j LOG --log-prefix "filt inp [ztunnel-rdhfd] "
-A FORWARD -j LOG --log-prefix "filt fw [ztunnel-rdhfd] "
-A OUTPUT -j LOG --log-prefix "filt out [ztunnel-rdhfd] "
COMMIT

*raw
:PREROUTING ACCEPT [13944:25638682]
:OUTPUT ACCEPT [8337:431154]
-A PREROUTING -j LOG --log-prefix "raw pre [ztunnel-rdhfd] "
-A OUTPUT -j LOG --log-prefix "raw out [ztunnel-rdhfd] "
COMMIT

*nat
:PREROUTING ACCEPT [200:18365]
:INPUT ACCEPT [129:10628]
:OUTPUT ACCEPT [102:9316]
:POSTROUTING ACCEPT [102:9316]
-A PREROUTING -j LOG --log-prefix "nat pre [ztunnel-rdhfd] "
-A INPUT -j LOG --log-prefix "nat inp [ztunnel-rdhfd] "
-A OUTPUT -j LOG --log-prefix "nat out [ztunnel-rdhfd] "
-A OUTPUT -p tcp -m tcp --dport 15088 -j REDIRECT --to-ports 15008
-A POSTROUTING -j LOG --log-prefix "nat post [ztunnel-rdhfd] "
COMMIT

*mangle
:PREROUTING ACCEPT [13872:25631773]
:INPUT ACCEPT [13872:25630905]
:FORWARD ACCEPT [0:0]
:OUTPUT ACCEPT [8337:431154]
:POSTROUTING ACCEPT [8337:431154]
-A PREROUTING -j LOG --log-prefix "mangle pre [ztunnel-rdhfd] "
-A PREROUTING -i pistioin -p tcp -m tcp --dport 15008 -j TPROXY --on-port 15008 --on-ip 127.0.0.1 --tproxy-mark 0x400/0xfff
-A PREROUTING -i pistioout -p tcp -j TPROXY --on-port 15001 --on-ip 127.0.0.1 --tproxy-mark 0x400/0xfff
-A PREROUTING -i pistioin -p tcp -j TPROXY --on-port 15006 --on-ip 127.0.0.1 --tproxy-mark 0x400/0xfff
-A PREROUTING ! -d 10.244.1.2/32 -i eth0 -p tcp -j MARK --set-xmark 0x4d3/0xfff
-A INPUT -j LOG --log-prefix "mangle inp [ztunnel-rdhfd] "
-A FORWARD -j LOG --log-prefix "mangle fw [ztunnel-rdhfd] "
-A OUTPUT -j LOG --log-prefix "mangle out [ztunnel-rdhfd] "
-A POSTROUTING -j LOG --log-prefix "mangle post [ztunnel-rdhfd] "
COMMIT

$ ip rule
0:      from all lookup local
20000:  from all fwmark 0x400/0xfff lookup 100
20001:  from all fwmark 0x401/0xfff lookup 101
20002:  from all fwmark 0x402/0xfff lookup 102
20003:  from all fwmark 0x4d3/0xfff lookup 100
32766:  from all lookup main
32767:  from all lookup default

$ ip route show table 100
local default dev lo scope host 

$ ip route show table 101
default via 192.168.127.1 dev pistioout 
10.244.1.1 dev eth0 scope link 

$ ip route show table 102
default via 192.168.126.1 dev pistioin 
10.244.1.1 dev eth0 scope link 

$ ip route show table 100
local default dev lo scope host 

$ ip route show table main
default via 10.244.1.1 dev eth0 
10.244.1.0/24 via 10.244.1.1 dev eth0 src 10.244.1.2 
10.244.1.1 dev eth0 scope link src 10.244.1.2 
192.168.126.0/30 dev pistioin proto kernel scope link src 192.168.126.2 
192.168.127.0/30 dev pistioout proto kernel scope link src 192.168.127.2 
```