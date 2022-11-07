NodeName=$1
kubectl label nodes "$NodeName" offMeshNodeType=main
