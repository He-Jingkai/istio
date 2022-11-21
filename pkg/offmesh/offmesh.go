package offmesh

type PUPair struct {
	CPUIp   string `yaml:"cpuNodeIP"`
	DPUIp   string `yaml:"dpuNodeIP"`
	CPUName string `yaml:"cpuNodeName"`
	DPUName string `yaml:"dpuNodeName"`
}

type PU struct {
	IP   string `yaml:"nodeIP"`
	Name string `yaml:"nodeName"`
}
type ClusterConfig struct {
	Pairs   []PUPair `yaml:"pairs"`
	Singles []PU     `yaml:"singles"`
}

type NodeInfo struct {
	IsSingleNode bool
	IsCPUNode    bool
	IsDPUNode    bool
	IsMyCPUNode  bool
	DPUIp        string
}

//pairs:
//- cpuNodeIP: 192.168.50.130
//dpuNodeIP: 192.168.50.131
//cpuNodeName: master
//dpuNodeName: master-dpu
//- cpuNodeIP: 192.168.50.133
//dpuNodeIP: 192.168.50.128
//cpuNodeName: worker1
//dpuNodeName: worker1-dpu

func ReadClusterConfigYaml(filePath string) ClusterConfig {
	//var clusterConf ClusterConfig
	//var err error
	//file, err := os.ReadFile(filePath)
	//if err != nil {
	//	klog.Errorf("read cluster conf yaml error: %v", err)
	//}
	//err = yaml.Unmarshal(file, &clusterConf)
	//if err != nil {
	//	klog.Errorf("unmarshal cluster conf yaml error: %v", err)
	//}
	//return clusterConf
	return ClusterConfig{
		Pairs: []PUPair{{
			CPUIp:   "192.168.50.130",
			DPUIp:   "192.168.50.131",
			CPUName: "master",
			DPUName: "master-dpu",
		}, {
			CPUIp:   "192.168.50.133",
			DPUIp:   "192.168.50.128",
			CPUName: "worker1",
			DPUName: "worker1-dpu",
		}},
	}
}
