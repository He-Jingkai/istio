# Connection reset by peer
## envoy log
```
2022-10-05T08:29:05.230956Z	debug	envoy filter	original_dst: new connection accepted
2022-10-05T08:29:05.230990Z	trace	envoy filter	original_dst: set destination to 10.32.0.11:15006
2022-10-05T08:29:05.231533Z	debug	envoy filter	[C2] new tcp proxy session
2022-10-05T08:29:05.231548Z	trace	envoy connection	[C2] readDisable: disable=true disable_count=0 state=0 buffer_length=0
2022-10-05T08:29:05.231573Z	debug	envoy filter	[C2] Creating connection to cluster BlackHoleCluster
2022-10-05T08:29:05.231594Z	debug	envoy upstream	no healthy host for TCP connection pool
2022-10-05T08:29:05.231598Z	debug	envoy connection	[C2] closing data_to_write=0 type=1
2022-10-05T08:29:05.231602Z	debug	envoy connection	[C2] closing socket: 1
2022-10-05T08:29:05.231647Z	trace	envoy connection	[C2] raising connection event 1
```

## how
see https://github.com/envoyproxy/envoy/issues/23414