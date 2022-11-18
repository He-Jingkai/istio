# ambient network notes

```shell
 kubectl create configmap offmesh-conf -n ambient --from-file=/home/hjk/offmesh-conf
```

```yaml
volumeMounts:
  - mountPath: /etc/offmesh-conf
    name: offmesh-conf
      
volumes:
  - name: offmesh-conf
    configMap:
      name: offmesh-conf
```

```shell
#cpu 
kubectl label nodes "$NodeName" offMeshNodeType=cpu
#dpu
kubectl label nodes "$NodeName" offMeshNodeType=dpu
```
## original

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

# Make sure that whatever is skipped is also skipped for returning packets.
# Input chain might be needed for things in host namespace that are skipped.
# Place the mark here after routing was done, not sure if conn-tracking will figure
# it out if I do it before, as NAT might change the connection tuple.
-A ztunnel-FORWARD -m mark --mark 0x220/0x220 -j CONNMARK --save-mark --nfmask 0x220 --ctmask 0x220
-A ztunnel-FORWARD -m mark --mark 0x210/0x210 -j CONNMARK --save-mark --nfmask 0x210 --ctmask 0x210
-A ztunnel-INPUT -m mark --mark 0x220/0x220 -j CONNMARK --save-mark --nfmask 0x220 --ctmask 0x220
-A ztunnel-INPUT -m mark --mark 0x210/0x210 -j CONNMARK --save-mark --nfmask 0x210 --ctmask 0x210
-A ztunnel-OUTPUT -s 10.244.0.1/32 -j MARK --set-xmark 0x220/0xffffffff
# Skip things that come from the tunnels, but don't apply the conn skip mark. If we have a skip mark, save it to conn mark.
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

```go
package constants
const (
	OutboundMask = "0x100"
	OutboundMark = OutboundMask + "/" + OutboundMask
	SkipMask     = "0x200"
	SkipMark     = SkipMask + "/" + SkipMask
	ConnSkipMask = "0x220"
	ConnSkipMark = ConnSkipMask + "/" + ConnSkipMask
	ProxyMask    = "0x210"
	ProxyMark    = ProxyMask + "/" + ProxyMask
	ProxyRetMask = "0x040"
	ProxyRetMark = ProxyRetMask + "/" + ProxyRetMask
	
	InboundTun  = "istioin"
	OutboundTun = "istioout"

	InboundTunIP         = "192.168.126.1"
	ZTunnelInboundTunIP  = "192.168.126.2"
	OutboundTunIP        = "192.168.127.1"
	ZTunnelOutboundTunIP = "192.168.127.2"
	TunPrefix            = 30

	ChainZTunnelPrerouting  = "ztunnel-PREROUTING"
	ChainZTunnelPostrouting = "ztunnel-POSTROUTING"
	ChainZTunnelInput       = "ztunnel-INPUT"
	ChainZTunnelOutput      = "ztunnel-OUTPUT"
	ChainZTunnelForward     = "ztunnel-FORWARD"

	ChainPrerouting  = "PREROUTING"
	ChainPostrouting = "POSTROUTING"
	ChainInput       = "INPUT"
	ChainOutput      = "OUTPUT"
	ChainForward     = "FORWARD"
)
// cni/pkg/ambient/constants/constants.go
```

## CPU Node
```bash
$ ip rule
0:      from all lookup local
100:    from all fwmark 0x200/0x200 goto 32766
101:    from all fwmark 0x100/0x100 lookup 101
102:    from all fwmark 0x40/0x40 lookup 102
32766:  from all lookup main
32767:  from all lookup default

$ ip route show table 101
default via 192.168.128.2 dev dputun 

$ ip route show table 102
default via 192.168.50.131 dev ens32 

$ ip route show table 100
10.244.0.7 via 192.168.126.2 dev istioin src 10.244.0.1 
10.244.0.8 via 192.168.126.2 dev istioin src 10.244.0.1 
```

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
-A ztunnel-PREROUTING -i dputun -j MARK --set-xmark 0x200/0x200
-A ztunnel-PREROUTING -i dputun -j RETURN
-A ztunnel-PREROUTING -p udp -m udp --dport 6081 -j RETURN
-A ztunnel-PREROUTING -m connmark --mark 0x220/0x220 -j MARK --set-xmark 0x200/0x200 
-A ztunnel-PREROUTING -m mark --mark 0x200/0x200 -j RETURN 
-A ztunnel-PREROUTING ! -i ens32 -m connmark --mark 0x210/0x210 -j MARK --set-xmark 0x40/0x40
-A ztunnel-PREROUTING -m mark --mark 0x40/0x40 -j RETURN
-A ztunnel-PREROUTING ! -s 10.244.0.3/32 -i ens32 --match-set ztunnel-pods-ips dst -j MARK --set-xmark 0x210/0x210
-A ztunnel-PREROUTING -m mark --mark 0x200/0x200 -j RETURN
-A ztunnel-PREROUTING -i ens32 --match-set ztunnel-pods-ips dst -j MARK --set-xmark 0x220/0x220
-A ztunnel-PREROUTING -p udp -j MARK --set-xmark 0x220/0x220
-A ztunnel-PREROUTING -m mark --mark 0x200/0x200 -j RETURN
-A ztunnel-PREROUTING -p tcp -m set --match-set ztunnel-pods-ips src -j MARK --set-xmark 0x100/0x100

*nat
-A PREROUTING -j ztunnel-PREROUTING
-A POSTROUTING -j ztunnel-POSTROUTING
-A ztunnel-POSTROUTING -m mark --mark 0x100/0x100 -j ACCEPT
-A ztunnel-PREROUTING -m mark --mark 0x100/0x100 -j ACCEPT
```

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
