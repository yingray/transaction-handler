FROM golang:1.11.5-alpine as compiler
ADD . /gopath/src/my-project
ENV GOPATH /gopath
WORKDIR /gopath/src/my-project
RUN go build server.go

FROM alpine:3.8
RUN apk add --update ca-certificates && rm -rf /var/cache/apk/* /tmp/*
COPY --from=compiler /gopath/src/my-project/server /go/bin/server
ENTRYPOINT /go/bin/server
