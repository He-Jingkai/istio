docker build -t offmesh-proxy-init .
docker run -it offmesh-proxy-init
docker ps -a
docker commit 59c4d3ec6095 hejingkai/offmesh-proxy-init
docker push hejingkai/offmesh-proxy-init
# hejingkai/offmesh-proxy-init