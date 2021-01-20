FROM alpine:3.9
MAINTAINER jin

ENV BUILDDIR=/bitcoinpay

RUN apk add --no-cache curl bash jq curl && \
    rm -rf /var/cache/apk/*

COPY ./build/ $BUILDDIR

ENTRYPOINT ["/bitcoinpay/launch"]

