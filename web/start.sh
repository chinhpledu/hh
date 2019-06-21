#!/usr/bin/env bash
#

######## install django ########################
#update-alternatives --install /usr/bin/python python /usr/bin/python2.7 1
#update-alternatives --install /usr/bin/python python /usr/bin/python3.7 2
#update-alternatives --set python /usr/bin/python3.7

#THEN
# install pip
# curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py

#mv /usr/bin/python /usr/bin/python.1.bk
#ln /usr/bin/python3 /usr/bin/python
# python  get-pip.py
#
# Cai django
# git clone https://github.com/django/django.git
# setup django environment
# pip install -e django/

# Sau buoc nay moi truong django da duoc cai dat. Test thu cau lenh django-admin xem co hoat dong khong
# django-admin startproject web  # tao ra thu muc project web
# cd web
# chay thu server
#python manage.py runserver 0:80
#Bay gio vao browser mo http://localhost:9003, se thay giao dien mac dinh cua trang web
