FROM ubuntu:latest

MAINTAINER 南川 ファム <fam.minamikawa@gmail.com>

ENV DEBIAN_FRONTEND noninteractive


ENV PYTHONIOENCODING=UTF-8


RUN apt-get update && apt-get install -y supervisor && apt-get -y upgrade -y && \
    apt-get install iputils-ping -y && \
    apt-get install curl -y  && \
    apt-get install -y git && \
    apt-get install -y golang && \
    apt-get install vim -y

ADD ./socket/supervisord.conf /etc/supervisor/conf.d/supervisord.conf


COPY ./socket/core /socket/core

WORKDIR /socket/core

COPY ./socket/start.sh /usr/local/start.sh
RUN chmod +x /usr/local/start.sh
RUN /usr/local/start.sh

ENTRYPOINT ["/usr/bin/supervisord"]
