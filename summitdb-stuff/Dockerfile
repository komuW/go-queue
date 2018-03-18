FROM alpine:latest

MAINTAINER Peter Teich <mail@pteich.xyz>

ADD ca-certificates.crt /etc/ssl/certs/

ENV SUMMITDB_VERSION 0.4.0

RUN mkdir -p /opt

RUN set -x \
    && apk update && apk add --no-cache --virtual .deps \
        openssl \
        dpkg \
        ca-certificates \
    && update-ca-certificates \
    && cd /opt \
    && dpkgArch="$(dpkg --print-architecture | awk -F- '{ print $NF }')" \
    && wget -O summitdb.tar.gz "https://github.com/tidwall/summitdb/releases/download/${SUMMITDB_VERSION}/summitdb-${SUMMITDB_VERSION}-linux-${dpkgArch}.tar.gz" \
    && tar xzvf summitdb.tar.gz \
    && rm -f summitdb.tar.gz \
    && mv summitdb-${SUMMITDB_VERSION}-linux-${dpkgArch} summitdb \
    && apk del .deps

RUN mkdir /data1
VOLUME ["/data1"]

EXPOSE 7481

WORKDIR /opt/summitdb

CMD ./summitdb-server -dir /data1 -loglevel debug
