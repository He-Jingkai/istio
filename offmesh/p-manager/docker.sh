docker build -t p_manager .
docker run -it -p 80:80 p_manager
docker ps -a
docker commit b9abc4d95e7d hejingkai/p_manager
docker push hejingkai/p_manager
# hejingkai/p_manager

