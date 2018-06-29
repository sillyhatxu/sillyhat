# golang

## chapters 1

### basic/hello
### basic/bigdigits
### basic/stacker
### basic/
### basic/

## project
### monitoring
系统监控


### message
消息组件


````
docker build -t golang.temporary .
docker tag golang.temporary:latest 111909622691.dkr.ecr.ap-southeast-1.amazonaws.com/golang.temporary:dt-1.9
docker push 111909622691.dkr.ecr.ap-southeast-1.amazonaws.com/golang.temporary:dt-1.9

Golang build
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
go build main.go
go clean

Build
Docker Hub
docker build -t xushikuan/sillyhat.golang.message .
docker push xushikuan/sillyhat.golang.message

AWS
docker build -t golang.message .
docker tag golang.message:latest 111909622691.dkr.ecr.ap-southeast-1.amazonaws.com/golang.message:dp-1.0
docker push 111909622691.dkr.ecr.ap-southeast-1.amazonaws.com/golang.message:dp-1.0
docker pull 111909622691.dkr.ecr.ap-southeast-1.amazonaws.com/golang.message:dp-1.0
````
Start
````
docker run -d -p 18001:18001 message
````