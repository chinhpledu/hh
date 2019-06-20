FROM ubuntu:latest

MAINTAINER Chinh PL <chinhpl@wealth-park.com>

ENV DEBIAN_FRONTEND noninteractive


ENV PYTHONIOENCODING=UTF-8


RUN apt-get update && apt-get install -y supervisor


ADD ./supervisord.conf /etc/supervisor/conf.d/supervisord.conf


COPY ./server/ /var/www/html

WORKDIR /var/www/html

ENTRYPOINT ["/usr/bin/supervisord"]
