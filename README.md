## cls-go
    请求财联社新闻及热门板块信息，通过钉钉机器人发送群消息

## 项目介绍
    本项目基于Go语言实现财联社新闻及每日热门板块的准实时消息推送

## 项目运行
    
    配置GOPROXY
    # Enable the go modules feature
    export GO111MODULE=on
    # Set the GOPROXY environment variable
    export GOPROXY=https://goproxy.io
  
    执行：
    1 go mod tidy
    2 go build -o ./bin/finance main.go 
    3 ./bin/finance
    或者
    1 go run main.go

## 配置Dockerfile
    具体内容见Dockerfile文件
    构建镜像docker build -t cls .
    运行容器docker run cls
