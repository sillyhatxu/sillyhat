FROM golang:1.10 AS builder

WORKDIR /go
##copy current dir all file into docker
##COPY . .
### copy golang main file to docker dir
COPY monitoring.go /go

## HTTPS Service
#RUN apt-get update
#RUN apt-get install -y ca-certificates#

##build linux version go file
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags monitoring -o build/monitoring -ldflags '-w' .

# start from scratch
FROM scratch
LABEL MAINTAINER="heixiushamao@gmail.com"

### set golang environment
#ENV JAVA_HOME              /usr/lib/jvm/java-8-openjdk-amd64
#ENV PATH                   $PATH:$JAVA_HOME/bin
#ENV JAVA_OPTIONS    -Xmx256m

### set bash environment dir
WORKDIR /go/monitoring

### copy golang main file to docker dir from local
#COPY monitoring /go/monitoring
### copy golang main file to docker dir from builder location
COPY --from=builder /go/build/monitoring /go/monitoring
### add ca certification for https
##ADD ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

#compile linux only
#ENV GOOS=linux

### set current time
#ENV TIME_ZONE              Asia/Singapore
#RUN echo "$TIME#_ZONE" > /etc/timezone
#RUN dpkg-reconfigure -f noninteractive tzdata


# tell how to run this container
#CMD ["./monitoring"]
ENTRYPOINT ["./monitoring"]

########  example  ########
#######FROM instrumentisto/glide as builder
#######WORKDIR /go/src/github.com/go-scratch-example
#########copy current dir all file into docker
#######COPY . .
#######RUN glide install
#######RUN CGO_ENABLED=0 GOOS=linux go build -a -tags sctratch-exampl -o build/go-scratch-example -ldflags ‘-w’ .
#######
#######FROM scratch
#######COPY --from=builder /go/src/github.com/go-scratch-example/build/go-scratch-example app
#######ENV PORT 3000
#######EXPOSE 3000
#######ENTRYPOINT ["/app"]