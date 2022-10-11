docker build -t offmesh-proxy-init .
docker run -it offmesh-proxy-init
docker ps -a
docker commit f1b05c2af61a hejingkai/offmesh-proxy-init
docker push hejingkai/offmesh-proxy-init
# hejingkai/offmesh-proxy-init