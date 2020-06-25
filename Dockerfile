FROM golang:latest as builder 
RUN mkdir -p /go/src/github.com/ppetko/kickstart
ADD ./* /go/src/github.com/ppetko/kickstart
WORKDIR /go/src/github.com/ppetko/kickstart
RUN go get && go test ./...
RUN go build -o main .

FROM centos:7.5.1804
EXPOSE 8080
RUN mkdir -p /kickstart 
WORKDIR /kickstart

COPY --from=builder /go/src/github.com/ppetko/kickstart/main /kickstart/
ADD ./ks.tmpl /kickstart/
ENTRYPOINT ["/kickstart/main"]

