FROM ubuntu:latest

MAINTAINER 南川 ファム <fam.minamikawa@gmail.com>

ENV DEBIAN_FRONTEND noninteractive


ENV PYTHONIOENCODING=UTF-8


RUN apt-get update && apt-get install -y supervisor && apt-get -y upgrade -y && \
    apt-get install curl -y  && \
    apt-get install -y git && \
    apt-get install -y golang && \
    apt-get install vim -y

ADD ./supervisord.conf /etc/supervisor/conf.d/supervisord.conf


COPY ./server/ /var/www/html

WORKDIR /var/www/html

COPY ./start.sh /start.sh
RUN chmod +x /start.sh
RUN /start.sh

ENTRYPOINT ["/usr/bin/supervisord"]
