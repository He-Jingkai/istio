PROXY_NAME=$1 #test
POD_IP=$2 #10.32.0.6
PROXY_IP=$3 #10.32.0.8
# STEP 1: 将需要转发的网络包mark

# STEP 2: 将被mark过的网络包路由到proxy