docker build -t offmesh-init .
docker run -it offmesh-init
docker ps -a
docker commit f1b05c2af61a hejingkai/offmesh-init
docker push hejingkai/offmesh-init
# hejingkai/offmesh-init