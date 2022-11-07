SideNodeName=$1
WorkerNodeName=$2
WorkerNodeGatewayIP=$3
kubectl label nodes "$SideNodeName" offMeshNodeType=dpu
sudo mkdir -p /var/node_config
echo "$WorkerNodeGatewayIP" | sudo tee /var/node_config/worker_node_gateway_ip
echo "$WorkerNodeName" | sudo tee /var/node_config/worker_node_name
