FROM ubuntu
RUN apt-get update
RUN apt-get install iptables -y
RUN apt-get install iproute2 -y
ADD route.sh /
ENTRYPOINT ["/bin/bash","/route.sh"]