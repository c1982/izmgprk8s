FROM golang:1.12.5-alpine3.9 AS builder

RUN apk add git gcc musl-dev 

RUN mkdir -p /go/src \
 && mkdir -p /go/bin \
 && mkdir -p /go/pkg

ENV PATH=$GOPATH/bin:$PATH
ENV GO111MODULE=on

ADD . /go/src/izmgprk8s

WORKDIR /go/src/izmgprk8s
RUN go mod vendor
RUN go build -o app

FROM alpine:3.7 as app

COPY --from=builder /go/src/izmgprk8s/app /usr/local/bin
CMD /usr/local/bin/app