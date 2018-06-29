FROM golang:1.10 AS builder

WORKDIR /go
COPY . /go

RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/xushikuan/sillyhat-golang-tool/sillyhat_database

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

FROM scratch
LABEL MAINTAINER="heixiushamao@gmail.com"
WORKDIR /go
COPY --from=builder /go /go
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./main"]