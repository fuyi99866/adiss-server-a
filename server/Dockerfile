FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR /go/src/adiss-server-a
COPY . .
RUN go env && go build -o server .

FROM alpine:latest
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /go/src/adiss-server-a
COPY --from=0 /go/src/adiss-server-a/server ./
COPY --from=0 /go/src/adiss-server-a/config.yaml ./
COPY --from=0 /go/src/adiss-server-a/resource ./resource

EXPOSE 8888

ENTRYPOINT ./server
