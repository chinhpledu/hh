FROM ubuntu:latest

MAINTAINER 南川 ファム <fam.minamikawa@gmail.com>

ENV DEBIAN_FRONTEND noninteractive


ENV PYTHONIOENCODING=UTF-8


RUN apt-get update && apt-get install -y supervisor && apt-get -y upgrade -y && \
    apt-get install iputils-ping -y && \
    apt-get install curl -y  && \
    apt-get install wget -y && \
    apt-get install -y git && \
    apt-get install vim -y

RUN wget https://dl.google.com/go/go1.12.6.linux-amd64.tar.gz
RUN tar -xvf go1.12.6.linux-amd64.tar.gz
RUN mv go /usr/local

RUN echo "export GOROOT=/usr/local/go" >> ~/.bashrc
RUN echo "export GOPATH=/socket/core" >> ~/.bashrc
RUN echo "export PATH=$PATH:/usr/local/go/bin:/socket/core/bin" >> ~/.bashrc 

ADD ./socket/supervisord.conf /etc/supervisor/conf.d/supervisord.conf


COPY ./socket/core /socket/core

WORKDIR /socket/core

COPY ./socket/start.sh /usr/local/start.sh
RUN chmod +x /usr/local/start.sh
RUN /usr/local/start.sh

ENTRYPOINT ["/usr/bin/supervisord"]
