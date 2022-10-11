FROM ubuntu
RUN apt-get update
RUN apt-get install iptables -y
ADD redirect.sh /
ENTRYPOINT ["/bin/bash","/redirect.sh"]