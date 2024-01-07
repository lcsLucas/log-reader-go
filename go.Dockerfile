FROM golang:1.20.3-alpine

ARG user=lucas
ARG uid=1000

WORKDIR /go/src/log-reader-go

RUN apk add --no-cache \
    bash

RUN adduser -D -u 1000 -s /bin/bash ${user} && \
    chown -R ${user} . && \
    chown -R ${user} /var/log/${ENV_PATH} && \
    find /var/log/${ENV_PATH} -type f | xargs -I{} chmod -v 644 {} && \
    find /var/log/${ENV_PATH} -type d | xargs -I{} chmod -v 755 {} && \
    find . -type f | xargs -I{} chmod -v 644 {} && \
    find . -type d | xargs -I{} chmod -v 755 {};

USER ${user}

ENV HOME /home/${user}