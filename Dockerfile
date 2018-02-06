FROM centos:7
ENV GOPATH /root/go
# RUN yum -y update
RUN yum install -y iproute initscripts openssh* wget vim nc telnet nmap mlocate tree gcc httpd php git java-1.8.0-openjdk epel-release python-dev

RUN rpm --import https://mirror.go-repo.io/centos/RPM-GPG-KEY-GO-REPO && curl -s https://mirror.go-repo.io/centos/go-repo.repo | tee /etc/yum.repos.d/go-repo.repo && yum install -y golang

RUN mkdir -p $GOPATH && cd $GOPATH && mkdir -p src/github.com/akshaysuresh4793 && cd src/github.com/akshaysuresh4793 && git clone https://github.com/akshaysuresh4793/zapposapi && cd zapposapi && go get -v && go build main.go restaurant.go menu.go menuitem.go location.go db.go json.go error.go response.go && cp main /usr/bin/server