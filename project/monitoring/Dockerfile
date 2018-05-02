# start from scratch
FROM scratch
LABEL MAINTAINER="heixiushamao@gmail.com"

### set golang environment
#ENV JAVA_HOME              /usr/lib/jvm/java-8-openjdk-amd64
#ENV PATH                   $PATH:$JAVA_HOME/bin
#ENV JAVA_OPTIONS    -Xmx256m

### set bash environment dir
WORKDIR /go/monitoring

### copy golang main file to docker dir
COPY monitoring /go/monitoring

#compile linux only
#ENV GOOS=linux

### set current time
#ENV TIME_ZONE              Asia/Singapore
#RUN echo "$TIME#_ZONE" > /etc/timezone
#RUN dpkg-reconfigure -f noninteractive tzdata

#build the binary with debug information removed
#RUN go build  -ldflags '-w -s' -a -installsuffix cgo -o monitoring

# tell how to run this container
CMD ["./monitoring"]