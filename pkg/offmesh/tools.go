package offmesh

func GetPair(nodeName string, nodeType string, offmeshCluster ClusterConfig) PU {
	//TODO:暂时不考虑single node的问题
	if nodeType == CPUNode {
		for _, pair := range offmeshCluster.Pairs {
			if pair.CPUName == nodeName {
				return PU{IP: pair.DPUIp, Name: pair.DPUName}
			}
		}
		return PU{}
	} else {
		for _, pair := range offmeshCluster.Pairs {
			if pair.DPUName == nodeName {
				return PU{IP: pair.CPUIp, Name: pair.CPUName}
			}
		}
		return PU{}
	}
}

func GetMyPair(nodeName string, offmeshCluster ClusterConfig) PU {
	//TODO:暂时不考虑single node的问题
	for _, pair := range offmeshCluster.Pairs {
		if pair.CPUName == nodeName {
			return PU{IP: pair.CPUIp, Name: pair.CPUName}
		}
		if pair.DPUName == nodeName {
			return PU{IP: pair.DPUIp, Name: pair.DPUName}
		}
	}
	return PU{}
}

func MyNodeType(NodeName string, offmeshCluster ClusterConfig) string {
	for _, pair := range offmeshCluster.Pairs {
		if pair.CPUName == NodeName {
			return CPUNode
		}
		if pair.DPUName == NodeName {
			return DPUNode
		}
	}
	return ""
}
