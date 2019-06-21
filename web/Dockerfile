FROM ubuntu:latest

MAINTAINER 南川 ファム <fam.minamikawa@gmail.com>

ENV DEBIAN_FRONTEND noninteractive


ENV PYTHONIOENCODING=UTF-8


RUN apt-get update && \
    apt-get install iputils-ping -y && \
    apt-get install curl -y  && \
    apt-get install build-essential libpq-dev libssl-dev openssl libffi-dev zlib1g-dev -y && \
    apt-get install python3-pip python3.7-dev -y && \
    apt-get install python3.7 -y && \
    apt-get install -y git && \
    apt-get install -y supervisor && \
    apt-get install vim -y

ADD ./web/supervisord.conf /etc/supervisor/conf.d/supervisord.conf


COPY ./web/core /web/core

WORKDIR /web/core

COPY ./web/start.sh /usr/local/start.sh
RUN chmod +x /usr/local/start.sh
RUN /usr/local/start.sh

ENTRYPOINT ["/usr/bin/supervisord"]
