# FROM python:3.7-alpine
FROM ubuntu:14.04
#
WORKDIR /app
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y python3 python3-pip nginx libsqlite3-0 dialog apt-utils libglib2.0-0
RUN pip3 install asyncio
RUN pip3 install flask
ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8
ENV DEBUG false
ENV FLASK_APP run.py
ENV FLASK_RUN_HOST 0.0.0.0
ENV PORT 5000

COPY . .
#RUN apk add   libsqlite3-0
RUN cp -r /usr/lib/x86_64-linux-gnu/libsqlite3.so.0 ./lib/ && \
  cp -r /usr/lib/x86_64-linux-gnu/libsqlite3.so.0.8.6 ./lib/ && \
  cp -r ./lib/lib* /usr/lib && \
  ln -s /usr/lib/libsqlite3.so.0.8.6  /usr/lib/libsqlite3.so && \
  ln -s /usr/lib/libicudata.so.53.1 /usr/lib/libicudata.so.53 && \
  ln -s /usr/lib/libicudata.so.53.1 /usr/lib/libicudata.so.5 && \
  ln -s /usr/lib/libicudata.so.53.1 /usr/lib/libicudata.so && \
  ln -s /usr/lib/libicui18n.so.53.1 /usr/lib/libicui18n.so.53 && \
  ln -s /usr/lib/libicui18n.so.53.1 /usr/lib/libicui18n.so.5 && \
  ln -s /usr/lib/libicui18n.so.53.1 /usr/lib/libicui18n.so && \
  ln -s /usr/lib/libicuuc.so.53.1 /usr/lib/libicuuc.so.53 && \
  ln -s /usr/lib/libicuuc.so.53.1 /usr/lib/libicuuc.so.5 && \
  ln -s /usr/lib/libicuuc.so.53.1 /usr/lib/libicuuc.so && \
  ln -s /usr/lib/libQt5Core.so.5.4.1 /usr/lib/libQt5Core.so.5.4 && \
  ln -s /usr/lib/libQt5Core.so.5.4.1 /usr/lib/libQt5Core.so.5 && \
  ln -s /usr/lib/libQt5Core.so.5.4.1 /usr/lib/libQt5Core.so && \
  ln -s /usr/lib/libQt5Network.so.5.4.1 /usr/lib/libQt5Network.so.5.4 && \
  ln -s /usr/lib/libQt5Network.so.5.4.1 /usr/lib/libQt5Network.so.5 && \
  ln -s /usr/lib/libQt5Network.so.5.4.1 /usr/lib/libQt5Network.so && \
  ln -s /usr/lib/libQt5Script.so.5.4.1 /usr/lib/libQt5Script.so.5.4 && \
  ln -s /usr/lib/libQt5Script.so.5.4.1 /usr/lib/libQt5Script.so.5 && \
  ln -s /usr/lib/libQt5Script.so.5.4.1 /usr/lib/libQt5Script.so && \
  ln -s /usr/lib/libQt5Sql.so.5.4.1 /usr/lib/libQt5Sql.so.5.4 && \
  ln -s /usr/lib/libQt5Sql.so.5.4.1 /usr/lib/libQt5Sql.so.5 && \
  ln -s /usr/lib/libQt5Sql.so.5.4.1 /usr/lib/libQt5Sql.so && \
  pip3 install -r app/requirements.txt --src /usr/local/src


# EXPOSE PORT
CMD ["flask", "run"]