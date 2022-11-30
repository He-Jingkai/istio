package offmesh

import (
	"gopkg.in/yaml.v2"
	"k8s.io/klog/v2"
	"os"
)

var offmeshCluster ClusterConfig
var read = false

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

func ReadClusterConfigYaml(filePath string) ClusterConfig {
	if read {
		return offmeshCluster
	}
	var err error
	file, err := os.ReadFile(filePath)
	if err != nil {
		klog.Errorf("read cluster conf yaml error: %v", err)
	}
	err = yaml.Unmarshal(file, &offmeshCluster)
	if err != nil {
		klog.Errorf("unmarshal cluster conf yaml error: %v", err)
	}
	read = true
	return offmeshCluster
	//return ClusterConfig{
	//	Pairs: []PUPair{{
	//		CPUIp:   "192.168.50.130",
	//		DPUIp:   "192.168.50.131",
	//		CPUName: "master",
	//		DPUName: "master-dpu",
	//	}, {
	//		CPUIp:   "192.168.50.133",
	//		DPUIp:   "192.168.50.128",
	//		CPUName: "worker1",
	//		DPUName: "worker1-dpu",
	//	}},
	//}
}
