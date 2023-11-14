FROM golang:alpine as build

ENV GOPROXY=https://proxy.golang.org

WORKDIR /go/src/adapter
COPY . .

RUN GOOS=linux go build -o /go/bin/adapter ./cmd/api/main.go

FROM alpine

COPY --from=build /go/bin/adapter /go/bin/adapter

RUN mkdir /root/.aws

COPY ~/.aws/credentials /root/.aws/credentials
COPY ~/.aws/config /root/.aws/config

ENTRYPOINT [ "/go/bin/adapter" ]