FROM golang:1.10 AS builder

WORKDIR /go
COPY . /go

#COPY /Users/xushikuan/go/gopath/src/github.com /go

RUN go get -u github.com/bsm/sarama-cluster
RUN go get -u github.com/Shopify/sarama
#RUN go get -u github.com/golang/glog
RUN go get -u github.com/gorilla/mux

# 如果你想在Linux 64位系统下运行
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build message.go
#WORKDIR /go
#COPY monitoring.go /go
#RUN CGO_ENABLED=0 GOOS=linux go build -a -tags monitoring -o build/monitoring -ldflags '-w' .
FROM scratch
LABEL MAINTAINER="heixiushamao@gmail.com"
WORKDIR /go
COPY --from=builder /go /go
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./message","-enviroment=dt"]
#ENTRYPOINT ["./message"]

