FROM ubuntu:latest

MAINTAINER 南川 ファム <fam.minamikawa@gmail.com>

ENV DEBIAN_FRONTEND noninteractive


ENV PYTHONIOENCODING=UTF-8


RUN apt-get update && \
    apt-get install iputils-ping -y && \
    apt-get install -y supervisor && \
    apt-get install vim -y

ADD ./web/supervisord.conf /etc/supervisor/conf.d/supervisord.conf


COPY ./web/core /web/core

WORKDIR /web/core

COPY ./web/start.sh /usr/local/start.sh
RUN chmod +x /usr/local/start.sh
RUN /usr/local/start.sh

ENTRYPOINT ["/usr/bin/supervisord"]
