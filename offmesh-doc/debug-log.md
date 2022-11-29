
[//]: # (abandoned)
# 一次正常的通信中的网络包流向
## CLIENT
```shell
# root@client-master-6b998fb777-29xzr:/apps# curl http://10.244.1.3
# eth0
# This is 10.244.0.2
14:05:13.611017 ARP, Request who-has 10.244.0.1 tell client-master-6b998fb777-29xzr, length 28
14:05:13.611066 ARP, Reply 10.244.0.1 is-at 62:dc:a4:c9:43:c6 (oui Unknown), length 28
14:05:13.711137 IP client-master-6b998fb777-29xzr.58408 > kube-dns.kube-system.svc.cluster.local.53: 25496+ PTR? 1.0.244.10.in-addr.arpa. (41)
14:05:13.715884 IP kube-dns.kube-system.svc.cluster.local.53 > client-master-6b998fb777-29xzr.58408: 25496 NXDomain 0/0/0 (41)
14:05:13.798360 IP client-master-6b998fb777-29xzr.50454 > kube-dns.kube-system.svc.cluster.local.53: 8775+ PTR? 10.0.96.10.in-addr.arpa. (41)
14:05:13.798984 IP kube-dns.kube-system.svc.cluster.local.53 > client-master-6b998fb777-29xzr.50454: 8775*- 1/0/0 PTR kube-dns.kube-system.svc.cluster.local. (116)
14:05:14.529285 IP client-master-6b998fb777-29xzr.40690 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.80: Flags [S], seq 3286940267, win 64240, options [mss 1460,sackOK,TS val 887486405 ecr 0,nop,wscale 7], length 0
14:05:14.529411 IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.80 > client-master-6b998fb777-29xzr.40690: Flags [S.], seq 95860346, ack 3286940268, win 64308, options [mss 1410,sackOK,TS val 3192380378 ecr 887486405,nop,wscale 7], length 0
14:05:14.529448 IP client-master-6b998fb777-29xzr.40690 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.80: Flags [.], ack 1, win 502, options [nop,nop,TS val 887486405 ecr 3192380378], length 0
14:05:14.529667 IP client-master-6b998fb777-29xzr.40690 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.80: Flags [P.], seq 1:75, ack 1, win 502, options [nop,nop,TS val 887486405 ecr 3192380378], length 74: HTTP: GET / HTTP/1.1
14:05:14.529719 IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.80 > client-master-6b998fb777-29xzr.40690: Flags [.], ack 75, win 502, options [nop,nop,TS val 3192380378 ecr 887486405], length 0
14:05:14.567942 IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.80 > client-master-6b998fb777-29xzr.40690: Flags [P.], seq 1:131, ack 75, win 502, options [nop,nop,TS val 3192380416 ecr 887486405], length 130: HTTP: HTTP/1.1 200 OK
14:05:14.567968 IP client-master-6b998fb777-29xzr.40690 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.80: Flags [.], ack 131, win 501, options [nop,nop,TS val 887486443 ecr 3192380416], length 0
14:05:14.569558 IP client-master-6b998fb777-29xzr.40690 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.80: Flags [F.], seq 75, ack 131, win 501, options [nop,nop,TS val 887486445 ecr 3192380416], length 0
14:05:14.575668 IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.80 > client-master-6b998fb777-29xzr.40690: Flags [F.], seq 131, ack 76, win 502, options [nop,nop,TS val 3192380424 ecr 887486445], length 0
14:05:14.575694 IP client-master-6b998fb777-29xzr.40690 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.80: Flags [.], ack 132, win 501, options [nop,nop,TS val 887486451 ecr 3192380424], length 0
14:05:14.730667 IP client-master-6b998fb777-29xzr.57435 > kube-dns.kube-system.svc.cluster.local.53: 18858+ PTR? 3.1.244.10.in-addr.arpa. (41)
14:05:14.731290 IP kube-dns.kube-system.svc.cluster.local.53 > client-master-6b998fb777-29xzr.57435: 18858*- 1/0/0 PTR 10-244-1-3.server-worker1-svc.default.svc.cluster.local. (133)
```
## CLIENT NODE'S ZTUNNEL
```shell
# pistioout
14:08:53.652357 IP 10.244.0.2.53932 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.http: Flags [S], seq 952168260, win 64240, options [mss 1460,sackOK,TS val 887705528 ecr 0,nop,wscale 7], length 0
14:08:53.652402 IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.http > 10.244.0.2.53932: Flags [S.], seq 3041611006, ack 952168261, win 64308, options [mss 1410,sackOK,TS val 3192599501 ecr 887705528,nop,wscale 7], length 0
14:08:53.652471 IP 10.244.0.2.53932 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.http: Flags [.], ack 1, win 502, options [nop,nop,TS val 887705528 ecr 3192599501], length 0
14:08:53.657714 IP 10.244.0.2.53932 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.http: Flags [P.], seq 1:75, ack 1, win 502, options [nop,nop,TS val 887705533 ecr 3192599501], length 74: HTTP: GET / HTTP/1.1
14:08:53.657746 IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.http > 10.244.0.2.53932: Flags [.], ack 75, win 502, options [nop,nop,TS val 3192599506 ecr 887705533], length 0
14:08:53.670683 IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.http > 10.244.0.2.53932: Flags [P.], seq 1:131, ack 75, win 502, options [nop,nop,TS val 3192599519 ecr 887705533], length 130: HTTP: HTTP/1.1 200 OK
14:08:53.670801 IP 10.244.0.2.53932 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.http: Flags [.], ack 131, win 501, options [nop,nop,TS val 887705546 ecr 3192599519], length 0
14:08:53.671249 IP 10.244.0.2.53932 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.http: Flags [F.], seq 75, ack 131, win 501, options [nop,nop,TS val 887705547 ecr 3192599519], length 0
14:08:53.715106 IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.http > 10.244.0.2.53932: Flags [.], ack 76, win 502, options [nop,nop,TS val 3192599564 ecr 887705547], length 0
14:08:53.787423 IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.http > 10.244.0.2.53932: Flags [F.], seq 131, ack 76, win 502, options [nop,nop,TS val 3192599636 ecr 887705547], length 0
14:08:53.787551 IP 10.244.0.2.53932 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.http: Flags [.], ack 132, win 501, options [nop,nop,TS val 887705663 ecr 3192599636], length 0
```

```shell
# eth0
14:11:08.116426 IP 10.244.0.2.60059 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.15008: Flags [S], seq 499675874, win 64240, options [mss 1460,sackOK,TS val 887839992 ecr 0,nop,wscale 7], length 0
14:11:08.118428 IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.15008 > 10.244.0.2.60059: Flags [S.], seq 1075512456, ack 499675875, win 64308, options [mss 1410,sackOK,TS val 3645624666 ecr 887839992,nop,wscale 7], length 0
14:11:08.118598 IP 10.244.0.2.60059 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.15008: Flags [.], ack 1, win 502, options [nop,nop,TS val 887839994 ecr 3645624666], length 0
14:11:08.120283 IP 10.244.0.1.28240 > ztunnel-tt74d.6081: Geneve, Flags [none], vni 0x3e9: IP 10.244.0.2.35888 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.http: Flags [P.], seq 1:75, ack 1, win 502, options [nop,nop,TS val 887839996 ecr 3192733964], length 74: HTTP: GET / HTTP/1.1
14:11:08.120495 IP ztunnel-tt74d.20615 > 10.244.0.1.6081: Geneve, Flags [none], vni 0x3e9: IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.http > 10.244.0.2.35888: Flags [.], ack 75, win 502, options [nop,nop,TS val 3192733969 ecr 887839996], length 0
14:11:08.121592 IP ztunnel-tt74d.58176 > kube-dns.kube-system.svc.cluster.local.domain: 3899+ PTR? 1.0.244.10.in-addr.arpa. (41)
14:11:08.126365 IP kube-dns.kube-system.svc.cluster.local.domain > ztunnel-tt74d.58176: 3899 NXDomain 0/0/0 (41)
14:11:08.128663 IP 10.244.0.2.60059 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.15008: Flags [P.], seq 1:2058, ack 1, win 502, options [nop,nop,TS val 887840004 ecr 3645624666], length 2057
14:11:08.129892 IP ztunnel-tt74d.37414 > kube-dns.kube-system.svc.cluster.local.domain: 39394+ PTR? 3.1.244.10.in-addr.arpa. (41)
14:11:08.130390 IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.15008 > 10.244.0.2.60059: Flags [.], ack 2058, win 496, options [nop,nop,TS val 3645624678 ecr 887840004], length 0
14:11:08.131032 IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.15008 > 10.244.0.2.60059: Flags [P.], seq 1:213, ack 2058, win 501, options [nop,nop,TS val 3645624678 ecr 887840004], length 212
14:11:08.131114 IP 10.244.0.2.60059 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.15008: Flags [.], ack 213, win 501, options [nop,nop,TS val 887840006 ecr 3645624678], length 0
14:11:08.132162 IP kube-dns.kube-system.svc.cluster.local.domain > ztunnel-tt74d.37414: 39394*- 1/0/0 PTR 10-244-1-3.server-worker1-svc.default.svc.cluster.local. (133)
14:11:08.170786 IP ztunnel-tt74d.40305 > kube-dns.kube-system.svc.cluster.local.domain: 29391+ PTR? 2.0.244.10.in-addr.arpa. (41)
14:11:08.177654 IP 10.244.0.2.60059 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.15008: Flags [P.], seq 2058:2122, ack 213, win 501, options [nop,nop,TS val 887840053 ecr 3645624678], length 64
14:11:08.179462 IP kube-dns.kube-system.svc.cluster.local.domain > ztunnel-tt74d.40305: 29391 NXDomain 0/0/0 (41)
14:11:08.183894 IP 10.244.0.2.60059 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.15008: Flags [P.], seq 2122:2281, ack 213, win 501, options [nop,nop,TS val 887840059 ecr 3645624678], length 159
14:11:08.185913 IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.15008 > 10.244.0.2.60059: Flags [.], ack 2281, win 501, options [nop,nop,TS val 3645624733 ecr 887840053], length 0
14:11:08.186818 IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.15008 > 10.244.0.2.60059: Flags [P.], seq 213:4047, ack 2281, win 501, options [nop,nop,TS val 3645624734 ecr 887840053], length 3834
14:11:08.186957 IP 10.244.0.2.60059 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.15008: Flags [.], ack 4047, win 488, options [nop,nop,TS val 887840062 ecr 3645624734], length 0
14:11:08.190447 IP 10.244.0.2.60059 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.15008: Flags [P.], seq 2281:2395, ack 4047, win 501, options [nop,nop,TS val 887840066 ecr 3645624734], length 114
14:11:08.194199 IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.15008 > 10.244.0.2.60059: Flags [P.], seq 4047:4208, ack 2395, win 501, options [nop,nop,TS val 3645624741 ecr 887840066], length 161
14:11:08.196431 IP ztunnel-tt74d.20615 > 10.244.0.1.6081: Geneve, Flags [none], vni 0x3e9: IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.http > 10.244.0.2.35888: Flags [P.], seq 1:131, ack 75, win 502, options [nop,nop,TS val 3192734045 ecr 887839996], length 130: HTTP: HTTP/1.1 200 OK
14:11:08.196555 IP 10.244.0.1.28240 > ztunnel-tt74d.6081: Geneve, Flags [none], vni 0x3e9: IP 10.244.0.2.35888 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.http: Flags [.], ack 131, win 501, options [nop,nop,TS val 887840072 ecr 3192734045], length 0
14:11:08.196993 IP 10.244.0.1.28240 > ztunnel-tt74d.6081: Geneve, Flags [none], vni 0x3e9: IP 10.244.0.2.35888 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.http: Flags [F.], seq 75, ack 131, win 501, options [nop,nop,TS val 887840072 ecr 3192734045], length 0
14:11:08.198207 IP 10.244.0.2.60059 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.15008: Flags [P.], seq 2395:2426, ack 4208, win 501, options [nop,nop,TS val 887840074 ecr 3645624741], length 31
14:11:08.201200 IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.15008 > 10.244.0.2.60059: Flags [P.], seq 4208:4239, ack 2426, win 501, options [nop,nop,TS val 3645624748 ecr 887840074], length 31
14:11:08.201498 IP ztunnel-tt74d.20615 > 10.244.0.1.6081: Geneve, Flags [none], vni 0x3e9: IP 10-244-1-3.server-worker1-svc.default.svc.cluster.local.http > 10.244.0.2.35888: Flags [F.], seq 131, ack 76, win 502, options [nop,nop,TS val 3192734050 ecr 887840072], length 0
14:11:08.201561 IP 10.244.0.1.28240 > ztunnel-tt74d.6081: Geneve, Flags [none], vni 0x3e9: IP 10.244.0.2.35888 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.http: Flags [.], ack 132, win 501, options [nop,nop,TS val 887840077 ecr 3192734050], length 0
14:11:08.239415 IP ztunnel-tt74d.44311 > kube-dns.kube-system.svc.cluster.local.domain: 11914+ PTR? 10.0.96.10.in-addr.arpa. (41)
14:11:08.240494 IP kube-dns.kube-system.svc.cluster.local.domain > ztunnel-tt74d.44311: 11914*- 1/0/0 PTR kube-dns.kube-system.svc.cluster.local. (116)
14:11:08.242987 IP 10.244.0.2.60059 > 10-244-1-3.server-worker1-svc.default.svc.cluster.local.15008: Flags [.], ack 4239, win 501, options [nop,nop,TS val 887840118 ecr 3645624748], length 0
14:11:08.939022 IP ztunnel-tt74d.39806 > istiod.istio-system.svc.cluster.local.15012: Flags [.], ack 2825447766, win 501, options [nop,nop,TS val 4269182164 ecr 623376036], length 0
14:11:08.940188 IP istiod.istio-system.svc.cluster.local.15012 > ztunnel-tt74d.39806: Flags [.], ack 1, win 501, options [nop,nop,TS val 623391124 ecr 4269167074], length 0
14:11:08.951571 IP ztunnel-tt74d.52159 > kube-dns.kube-system.svc.cluster.local.domain: 37081+ PTR? 242.44.103.10.in-addr.arpa. (44)
14:11:08.952670 IP kube-dns.kube-system.svc.cluster.local.domain > ztunnel-tt74d.52159: 37081*- 1/0/0 PTR istiod.istio-system.svc.cluster.local. (121)
14:11:08.985872 IP istiod.istio-system.svc.cluster.local.15012 > ztunnel-tt74d.39806: Flags [.], ack 1, win 501, options [nop,nop,TS val 623391170 ecr 4269167074], length 0
14:11:08.985903 IP ztunnel-tt74d.39806 > istiod.istio-system.svc.cluster.local.15012: Flags [.], ack 1, win 501, options [nop,nop,TS val 4269182211 ecr 623391124], length 0
```

```shell
#log
[2022-11-28T14:11:08.116Z] "- - -" 0 - - - "-" 74 130 85 - "-" "-" "-" "-" "10.244.1.3:15008" outbound_pod_tunnel_clus_spiffe://cluster.local/ns/default/sa/default 10.244.0.2:60059 10.244.1.3:80 10.244.0.2:35888 - - outbound capture listener
[2022-11-28T14:11:08.116Z] "- - -" 0 - - - "-" 74 130 85 - "-" "-" "-" "-" "10.244.1.3:15008" outbound_pod_tunnel_clus_spiffe://cluster.local/ns/default/sa/default 10.244.0.2:60059 10.244.1.3:80 10.244.0.2:35888 - - capture outbound pod (no waypoint proxy)
```
## SERVER NODE'S ZTUNNEL
```shell
#log
[2022-11-28T14:15:17.879Z] "CONNECT - HTTP/2" 200 - via_upstream - "-" 74 130 62 - "-" "-" "4ade89f8-cd47-4c66-b831-624f0a3b643a" "10.244.1.3:80" "10.244.1.3:80" virtual_inbound 10.244.0.2:51511 10.244.1.3:15008 10.244.0.2:54987 - - inbound hcm
```
## SERVER
```shell

```