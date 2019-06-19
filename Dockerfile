FROM ubuntu:18.04

MAINTAINER Chinh PL <fam.minamikawa@gmail.com>

COPY ./server /var/www/html

WORKDIR /var/www/html
