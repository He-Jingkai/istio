## first edition
```bash
*mangle
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
-A ztunnel-OUTPUT -s 10.244.0.1/32 -j MARK --set-xmark 0x220/0xffffffff
-A ztunnel-PREROUTING -i dputunnel -j MARK --set-xmark 0x200/0x200
-A ztunnel-PREROUTING -i dputunnel -j RETURN
-A ztunnel-PREROUTING -p udp -m udp --dport 6081 -j RETURN
-A ztunnel-PREROUTING -m connmark --mark 0x220/0x220 -j MARK --set-xmark 0x200/0x200
-A ztunnel-PREROUTING -m mark --mark 0x200/0x200 -j RETURN
-A ztunnel-PREROUTING ! -i ens32 -m connmark --mark 0x210/0x210 -j MARK --set-xmark 0x40/0x40
-A ztunnel-PREROUTING -m mark --mark 0x40/0x40 -j RETURN
-A ztunnel-PREROUTING ! -s 10.244.1.3/32 -i ens32 -m set --match-set ztunnel-pods-ips dst -j MARK --set-xmark 0x210/0x210
-A ztunnel-PREROUTING -m mark --mark 0x200/0x200 -j RETURN
-A ztunnel-PREROUTING -i ens32 -m set --match-set ztunnel-pods-ips dst -j MARK --set-xmark 0x220/0x220
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

-A ztunnel-POSTROUTING -m mark --mark 0x100/0x100 -j ACCEPT
-A ztunnel-PREROUTING -m mark --mark 0x100/0x100 -j ACCEPT
COMMIT

$ ip rule
0:      from all lookup local
98:     from all lookup 105
100:    from all fwmark 0x200/0x200 goto 32766
101:    from all fwmark 0x100/0x100 lookup 101
102:    from all fwmark 0x40/0x40 lookup 102
103:    from all lookup 100
32766:  from all lookup main
32767:  from all lookup default

$ ip route show table 105
192.168.50.131 dev ens32 scope link 

$ ip route show table 101
default via 192.168.128.2 dev dputunnel 

$ ip route show table 102
default via 192.168.50.131 dev ens32 

$ ip route show table 100
Error: ipv4: FIB table does not exist.
Dump terminated

$ ip route show table main
default via 192.168.50.2 dev ens32 proto dhcp metric 100 
10.244.0.2 dev veth5b04e3df scope host 
10.244.0.3 dev veth89c0fd0c scope host 
10.244.0.4 dev vethe836a4f3 scope host 
10.244.0.6 dev vetha5633b2e scope host 
10.244.0.7 dev vethcba0679c scope host 
10.244.1.0/24 via 192.168.50.131 dev ens32 
10.244.2.0/24 via 192.168.50.131 dev ens32 
10.244.3.0/24 via 192.168.50.131 dev ens32 
192.168.50.0/24 dev ens32 proto kernel scope link src 192.168.50.130 metric 100 
192.168.128.0/30 dev dputunnel proto kernel scope link src 192.168.128.1 
```

## second edition
```shell
$ iptable-save
*mangle
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
-A ztunnel-INPUT -m mark --mark 0x220/0x220 -j CONNMARK --save-mark --nfmask 0x220 --ctmask 0x220
-A ztunnel-OUTPUT -s 10.244.0.1/32 -j MARK --set-xmark 0x220/0xffffffff
-A ztunnel-PREROUTING -m connmark --mark 0x220/0x220 -j MARK --set-xmark 0x200/0x200
-A ztunnel-PREROUTING -m mark --mark 0x200/0x200 -j RETURN
-A ztunnel-PREROUTING -i ens32 -m set --match-set ztunnel-pods-ips dst -j MARK --set-xmark 0x200/0x200
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

-A ztunnel-POSTROUTING -m mark --mark 0x100/0x100 -j ACCEPT
-A ztunnel-PREROUTING -m mark --mark 0x100/0x100 -j ACCEPT
COMMIT
```

```shell
$ ip rule
0:      from all lookup local
100:    from all fwmark 0x200/0x200 goto 32766
101:    from all fwmark 0x100/0x100 lookup 101
32766:  from all lookup main
32767:  from all lookup default

$ ip route show table 101
default via 192.168.50.131 dev ens32 

$ ip route show table main
default via 192.168.50.2 dev ens32 proto dhcp metric 100 
10.244.0.2 dev veth5b04e3df scope host 
10.244.0.3 dev veth89c0fd0c scope host 
10.244.0.4 dev vethe836a4f3 scope host 
10.244.0.6 dev vetha5633b2e scope host 
10.244.0.7 dev vethcba0679c scope host 
10.244.1.0/24 via 192.168.50.131 dev ens32 
10.244.2.0/24 via 192.168.50.131 dev ens32 
10.244.3.0/24 via 192.168.50.131 dev ens32 
192.168.50.0/24 dev ens32 proto kernel scope link src 192.168.50.130 metric 100
```