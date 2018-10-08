FROM golang:1.11-alpine

ENV GOBIN=/go/bin

ADD . /go/github/giantswarm/semver-bump/.
RUN apk add -U --no-cache git && \
    cd /go/github/giantswarm/semver-bump && \
    go get && \
    go install && \
    apk del git

ENTRYPOINT [ "/go/bin/semver-bump" ]
