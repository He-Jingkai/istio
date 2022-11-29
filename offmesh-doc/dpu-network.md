## first edition

```shell
$ sudo iptables-save

*mangle
-A PREROUTING -j ztunnel-PREROUTING
-A PREROUTING -j ztunnel-PREROUTING
-A INPUT -j ztunnel-INPUT
-A INPUT -j ztunnel-INPUT
-A FORWARD -j ztunnel-FORWARD
-A FORWARD -j ztunnel-FORWARD
-A OUTPUT -j ztunnel-OUTPUT
-A OUTPUT -j ztunnel-OUTPUT
-A POSTROUTING -j ztunnel-POSTROUTING
-A POSTROUTING -j ztunnel-POSTROUTING
-A ztunnel-FORWARD -m mark --mark 0x220/0x220 -j CONNMARK --save-mark --nfmask 0x220 --ctmask 0x220
-A ztunnel-FORWARD -m mark --mark 0x210/0x210 -j CONNMARK --save-mark --nfmask 0x210 --ctmask 0x210
-A ztunnel-INPUT -m mark --mark 0x220/0x220 -j CONNMARK --save-mark --nfmask 0x220 --ctmask 0x220
-A ztunnel-INPUT -m mark --mark 0x210/0x210 -j CONNMARK --save-mark --nfmask 0x210 --ctmask 0x210
-A ztunnel-OUTPUT -s 10.244.1.1/32 -j MARK --set-xmark 0x220/0xffffffff
-A ztunnel-PREROUTING -i istioin -j MARK --set-xmark 0x240/0x240
-A ztunnel-PREROUTING -i istioin -j RETURN
-A ztunnel-PREROUTING -i istioout -j MARK --set-xmark 0x240/0x240
-A ztunnel-PREROUTING -i istioout -j RETURN
-A ztunnel-PREROUTING -p udp -m udp --dport 6081 -j RETURN
-A ztunnel-PREROUTING -m connmark --mark 0x220/0x220 -j MARK --set-xmark 0x200/0x200
-A ztunnel-PREROUTING -m mark --mark 0x200/0x200 -j RETURN
-A ztunnel-PREROUTING ! -i veth9f0ba668 -m connmark --mark 0x210/0x210 -j MARK --set-xmark 0x40/0x40
-A ztunnel-PREROUTING -m mark --mark 0x40/0x40 -j RETURN
-A ztunnel-PREROUTING ! -s 10.244.1.3/32 -i veth9f0ba668 -j MARK --set-xmark 0x210/0x210
-A ztunnel-PREROUTING -m mark --mark 0x200/0x200 -j RETURN
-A ztunnel-PREROUTING -i veth9f0ba668 -j MARK --set-xmark 0x220/0x220
-A ztunnel-PREROUTING -p udp -j MARK --set-xmark 0x220/0x220
-A ztunnel-PREROUTING -m mark --mark 0x200/0x200 -j RETURN
-A ztunnel-PREROUTING -p tcp -m set --match-set ztunnel-pods-ips src -j MARK --set-xmark 0x100/0x100
COMMIT

*nat
-A PREROUTING -j ztunnel-PREROUTING
-A PREROUTING -j ztunnel-PREROUTING
-A PREROUTING -m comment --comment "kubernetes service portals" -j KUBE-SERVICES
-A OUTPUT -m comment --comment "kubernetes service portals" -j KUBE-SERVICES
-A POSTROUTING -j ztunnel-POSTROUTING
-A POSTROUTING -j ztunnel-POSTROUTING
-A ztunnel-POSTROUTING -m mark --mark 0x100/0x100 -j ACCEPT
-A ztunnel-PREROUTING -m mark --mark 0x100/0x100 -j ACCEPT
COMMIT

$ ip rule
0:      from all lookup local
98:     from all lookup 105
99:     from all fwmark 0x240/0x240 lookup 104
100:    from all fwmark 0x200/0x200 goto 32766
101:    from all fwmark 0x100/0x100 lookup 101
102:    from all fwmark 0x40/0x40 lookup 102
103:    from all lookup 100
32766:  from all lookup main
32767:  from all lookup default

$ ip route show table local
local 10.244.1.1 dev veth9f0ba668 proto kernel scope host src 10.244.1.1 
broadcast 127.0.0.0 dev lo proto kernel scope link src 127.0.0.1 
local 127.0.0.0/8 dev lo proto kernel scope host src 127.0.0.1 
local 127.0.0.1 dev lo proto kernel scope host src 127.0.0.1 
broadcast 127.255.255.255 dev lo proto kernel scope link src 127.0.0.1 
broadcast 192.168.50.0 dev ens32 proto kernel scope link src 192.168.50.131 
local 192.168.50.131 dev ens32 proto kernel scope host src 192.168.50.131 
broadcast 192.168.50.255 dev ens32 proto kernel scope link src 192.168.50.131 
broadcast 192.168.126.0 dev istioin proto kernel scope link src 192.168.126.1 
local 192.168.126.1 dev istioin proto kernel scope host src 192.168.126.1 
broadcast 192.168.126.3 dev istioin proto kernel scope link src 192.168.126.1 
broadcast 192.168.127.0 dev istioout proto kernel scope link src 192.168.127.1 
local 192.168.127.1 dev istioout proto kernel scope host src 192.168.127.1 
broadcast 192.168.127.3 dev istioout proto kernel scope link src 192.168.127.1 
broadcast 192.168.128.0 dev cputunnel proto kernel scope link src 192.168.128.2 
local 192.168.128.2 dev cputunnel proto kernel scope host src 192.168.128.2 
broadcast 192.168.128.3 dev cputunnel proto kernel scope link src 192.168.128.2 

$ ip route show table 105
192.168.50.130 dev ens32 scope link 

$ ip route show table 104
default via 192.168.128.1 dev cputunnel 

$ ip route show table 101
default via 192.168.127.2 dev istioout 
10.244.1.3 dev veth9f0ba668 scope link 

$ ip route show table 102
default via 10.244.1.3 dev veth9f0ba668 onlink 
10.244.1.3 dev veth9f0ba668 scope link 

$ ip route show table 100
10.244.0.2 via 192.168.126.2 dev istioin src 10.244.1.1 
10.244.0.7 via 192.168.126.2 dev istioin src 10.244.1.1 
10.244.1.3 dev veth9f0ba668 scope link 

$ ip route show table main
default via 192.168.50.2 dev ens32 proto dhcp metric 100 
10.244.0.0/24 via 192.168.50.130 dev ens32 
10.244.1.3 dev veth9f0ba668 scope host 
10.244.2.0/24 via 192.168.50.128 dev ens32 
10.244.3.0/24 via 192.168.50.128 dev ens32 
192.168.50.0/24 dev ens32 proto kernel scope link src 192.168.50.131 metric 100 
192.168.126.0/30 dev istioin proto kernel scope link src 192.168.126.1 
192.168.127.0/30 dev istioout proto kernel scope link src 192.168.127.1 
192.168.128.0/30 dev cputunnel proto kernel scope link src 192.168.128.2 
```

## second edition
```shell
$ sudo iptables-save
*mangle
:PREROUTING ACCEPT [0:0]
:INPUT ACCEPT [0:0]
:FORWARD ACCEPT [0:0]
:OUTPUT ACCEPT [0:0]
:POSTROUTING ACCEPT [0:0]
:KUBE-IPTABLES-HINT - [0:0]
:KUBE-KUBELET-CANARY - [0:0]
:KUBE-PROXY-CANARY - [0:0]
:ztunnel-FORWARD - [0:0]
:ztunnel-INPUT - [0:0]
:ztunnel-OUTPUT - [0:0]
:ztunnel-POSTROUTING - [0:0]
:ztunnel-PREROUTING - [0:0]
-A PREROUTING -j ztunnel-PREROUTING
-A INPUT -j ztunnel-INPUT
-A FORWARD -j ztunnel-FORWARD
-A OUTPUT -j ztunnel-OUTPUT
-A POSTROUTING -j ztunnel-POSTROUTING
-A ztunnel-FORWARD -m mark --mark 0x220/0x220 -j CONNMARK --save-mark --nfmask 0x220 --ctmask 0x220
-A ztunnel-FORWARD -m mark --mark 0x210/0x210 -j CONNMARK --save-mark --nfmask 0x210 --ctmask 0x210
-A ztunnel-INPUT -m mark --mark 0x220/0x220 -j CONNMARK --save-mark --nfmask 0x220 --ctmask 0x220
-A ztunnel-INPUT -m mark --mark 0x210/0x210 -j CONNMARK --save-mark --nfmask 0x210 --ctmask 0x210
-A ztunnel-OUTPUT -s 10.244.1.1/32 -j MARK --set-xmark 0x220/0xffffffff
-A ztunnel-PREROUTING -i istioin -j MARK --set-xmark 0x200/0x200
-A ztunnel-PREROUTING -i istioin -j RETURN
-A ztunnel-PREROUTING -i istioout -j MARK --set-xmark 0x200/0x200
-A ztunnel-PREROUTING -i istioout -j RETURN
-A ztunnel-PREROUTING -p udp -m udp --dport 6081 -j RETURN
-A ztunnel-PREROUTING -m connmark --mark 0x220/0x220 -j MARK --set-xmark 0x200/0x200
-A ztunnel-PREROUTING -m mark --mark 0x200/0x200 -j RETURN
-A ztunnel-PREROUTING ! -i veth1ae58883 -m connmark --mark 0x210/0x210 -j MARK --set-xmark 0x40/0x40
-A ztunnel-PREROUTING -m mark --mark 0x40/0x40 -j RETURN
-A ztunnel-PREROUTING ! -s 10.244.1.2/32 -i veth1ae58883 -j MARK --set-xmark 0x210/0x210
-A ztunnel-PREROUTING -m mark --mark 0x200/0x200 -j RETURN
-A ztunnel-PREROUTING -i veth1ae58883 -j MARK --set-xmark 0x220/0x220
-A ztunnel-PREROUTING -p udp -j MARK --set-xmark 0x220/0x220
-A ztunnel-PREROUTING -m mark --mark 0x200/0x200 -j RETURN
-A ztunnel-PREROUTING -p tcp -m set --match-set ztunnel-pods-ips src -j MARK --set-xmark 0x100/0x100
COMMIT

*nat
:ztunnel-POSTROUTING - [0:0]
:ztunnel-PREROUTING - [0:0]
-A PREROUTING -j ztunnel-PREROUTING
-A PREROUTING -m comment --comment "kubernetes service portals" -j KUBE-SERVICES
-A OUTPUT -m comment --comment "kubernetes service portals" -j KUBE-SERVICES
-A POSTROUTING -j ztunnel-POSTROUTING
-A POSTROUTING -m comment --comment "kubernetes postrouting rules" -j KUBE-POSTROUTING
-A POSTROUTING -m addrtype ! --dst-type LOCAL -m comment --comment "kind-masq-agent: ensure nat POSTROUTING directs all non-LOCAL destination traffic to our custom KIND-MASQ-AGENT chain" -j KIND-MASQ-AGENT

-A ztunnel-POSTROUTING -m mark --mark 0x100/0x100 -j ACCEPT
-A ztunnel-PREROUTING -m mark --mark 0x100/0x100 -j ACCEPT
COMMIT

$ ip rule
0:      from all lookup local
100:    from all fwmark 0x200/0x200 goto 32766
101:    from all fwmark 0x100/0x100 lookup 101
102:    from all fwmark 0x40/0x40 lookup 102
103:    from all lookup 100
32766:  from all lookup main
32767:  from all lookup default

$ ip route show table 101
default via 192.168.127.2 dev istioout 
10.244.1.2 dev veth1ae58883 scope link 

$ ip route show table 102
default via 10.244.1.2 dev veth1ae58883 onlink 
10.244.1.2 dev veth1ae58883 scope link 

$ ip route show table 100
10.244.0.2 via 192.168.126.2 dev istioin src 10.244.1.1 
10.244.0.7 via 192.168.126.2 dev istioin src 10.244.1.1 
10.244.1.2 dev veth1ae58883 scope link 

$ ip route show table main
default via 192.168.50.2 dev ens32 proto dhcp metric 100 
10.244.0.0/24 via 192.168.50.130 dev ens32 
10.244.1.2 dev veth1ae58883 scope host 
10.244.2.0/24 via 192.168.50.128 dev ens32 
10.244.3.0/24 via 192.168.50.128 dev ens32 
192.168.50.0/24 dev ens32 proto kernel scope link src 192.168.50.131 metric 100 
192.168.126.0/30 dev istioin proto kernel scope link src 192.168.126.1 
192.168.127.0/30 dev istioout proto kernel scope link src 192.168.127.1 
```