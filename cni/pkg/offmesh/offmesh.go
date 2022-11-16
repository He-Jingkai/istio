package offmesh

import (
	"gopkg.in/yaml.v2"
	"k8s.io/klog/v2"
	"os"
)

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

const ClusterConfigYamlPath = `/etc/offmesh/cluster-conf.yaml`

func ReadClusterConfigYaml(filePath string) ClusterConfig {
	var clusterConf ClusterConfig
	var err error
	file, err := os.ReadFile(filePath)
	if err != nil {
		klog.Errorf("read cluster conf yaml error: %v", err)
	}
	err = yaml.Unmarshal(file, &clusterConf)
	if err != nil {
		klog.Errorf("unmarshal cluster conf yaml error: %v", err)
	}
	return clusterConf
}
