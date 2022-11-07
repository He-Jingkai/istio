# ambient network notes
## route
- `fwmark 0x200/0x200`: bypass policy routing
- `fwmark 0x40/0x40`, table 102: to ztunnel directly
- `fwmark 0x100/0x100`, table 101: ambient inbound
- `fwmark 0x210/0x210`:
- `fwmark 0x220/0x220`: ztunnel pod发出的网络包, 最终0x220会被标记到连接上，之后数据包会被标记为0x200
- 
```bash
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
10.244.0.3 dev vethfe50c7fb scope link 

$ ip route show table 102
default via 10.244.0.3 dev vethfe50c7fb onlink 
10.244.0.3 dev vethfe50c7fb scope link

$ ip route show table 100
10.244.0.3 dev vethfe50c7fb scope link 
10.244.0.7 via 192.168.126.2 dev istioin src 10.244.0.1 
10.244.0.8 via 192.168.126.2 dev istioin src 10.244.0.1 
```
## iptables
### by table
```bash
*mangle
# -A PREROUTING -j ztunnel-PREROUTING
-A INPUT -j ztunnel-INPUT
-A FORWARD -j ztunnel-FORWARD
-A OUTPUT -j ztunnel-OUTPUT
-A POSTROUTING -j ztunnel-POSTROUTING

-A ztunnel-FORWARD -m mark --mark 0x220/0x220 -j CONNMARK --save-mark --nfmask 0x220 --ctmask 0x220
-A ztunnel-FORWARD -m mark --mark 0x210/0x210 -j CONNMARK --save-mark --nfmask 0x210 --ctmask 0x210
-A ztunnel-INPUT -m mark --mark 0x220/0x220 -j CONNMARK --save-mark --nfmask 0x220 --ctmask 0x220
-A ztunnel-INPUT -m mark --mark 0x210/0x210 -j CONNMARK --save-mark --nfmask 0x210 --ctmask 0x210
-A ztunnel-OUTPUT -s 10.244.0.1/32 -j MARK --set-xmark 0x220/0xffffffff
-A ztunnel-PREROUTING -i istioin -j MARK --set-xmark 0x200/0x200
-A ztunnel-PREROUTING -i istioin -j RETURN
-A ztunnel-PREROUTING -i istioout -j MARK --set-xmark 0x200/0x200
-A ztunnel-PREROUTING -i istioout -j RETURN
-A ztunnel-PREROUTING -p udp -m udp --dport 6081 -j RETURN
-A ztunnel-PREROUTING -m connmark --mark 0x220/0x220 -j MARK --set-xmark 0x200/0x200
-A ztunnel-PREROUTING -m mark --mark 0x200/0x200 -j RETURN
-A ztunnel-PREROUTING ! -i vethfe50c7fb -m connmark --mark 0x210/0x210 -j MARK --set-xmark 0x40/0x40
-A ztunnel-PREROUTING -m mark --mark 0x40/0x40 -j RETURN
-A ztunnel-PREROUTING ! -s 10.244.0.3/32 -i vethfe50c7fb -j MARK --set-xmark 0x210/0x210
-A ztunnel-PREROUTING -m mark --mark 0x200/0x200 -j RETURN
-A ztunnel-PREROUTING -i vethfe50c7fb -j MARK --set-xmark 0x220/0x220
-A ztunnel-PREROUTING -p udp -j MARK --set-xmark 0x220/0x220
-A ztunnel-PREROUTING -m mark --mark 0x200/0x200 -j RETURN
-A ztunnel-PREROUTING -p tcp -m set --match-set ztunnel-pods-ips src -j MARK --set-xmark 0x100/0x100

*nat
-A PREROUTING -j ztunnel-PREROUTING
-A POSTROUTING -j ztunnel-POSTROUTING
-A ztunnel-POSTROUTING -m mark --mark 0x100/0x100 -j ACCEPT
-A ztunnel-PREROUTING -m mark --mark 0x100/0x100 -j ACCEPT
```
### by chain
#### PREROUTING
```bash
# mangle
-A PREROUTING -j ztunnel-PREROUTING

-A ztunnel-PREROUTING -i istioin -j MARK --set-xmark 0x200/0x200 #istioin's network traffic, no change
-A ztunnel-PREROUTING -i istioin -j RETURN #istioin's network traffic, no change
-A ztunnel-PREROUTING -i istioout -j MARK --set-xmark 0x200/0x200 #istioout's network traffic, no change
-A ztunnel-PREROUTING -i istioout -j RETURN #istioout's network traffic, no change
-A ztunnel-PREROUTING -p udp -m udp --dport 6081 -j RETURN #geneve type traffic traffic, no change
-A ztunnel-PREROUTING -m connmark --mark 0x220/0x220 -j MARK --set-xmark 0x200/0x200
-A ztunnel-PREROUTING -m mark --mark 0x200/0x200 -j RETURN
-A ztunnel-PREROUTING ! -i vethfe50c7fb -m connmark --mark 0x210/0x210 -j MARK --set-xmark 0x40/0x40
-A ztunnel-PREROUTING -m mark --mark 0x40/0x40 -j RETURN
-A ztunnel-PREROUTING ! -s 10.244.0.3/32 -i vethfe50c7fb -j MARK --set-xmark 0x210/0x210
-A ztunnel-PREROUTING -m mark --mark 0x200/0x200 -j RETURN
-A ztunnel-PREROUTING -i vethfe50c7fb -j MARK --set-xmark 0x220/0x220
-A ztunnel-PREROUTING -p udp -j MARK --set-xmark 0x220/0x220
-A ztunnel-PREROUTING -m mark --mark 0x200/0x200 -j RETURN
-A ztunnel-PREROUTING -p tcp -m set --match-set ztunnel-pods-ips src -j MARK --set-xmark 0x100/0x100

# nat
-A ztunnel-POSTROUTING -m mark --mark 0x100/0x100 -j ACCEPT
```

