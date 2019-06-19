FROM ubuntu:18.04

MAINTAINER Chinh PL <fam.minamikawa@gmail.com>

COPY ./src /var/www/html

WORKDIR /var/www/html
