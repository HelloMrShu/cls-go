FROM golang:1.16
ENV TZ=Asia/Shanghai
ENV GO111MODULE=on
ENV GOOS=linux
ENV	GOPROXY=https://goproxy.cn

WORKDIR /opt/projects/cls
COPY . /opt/projects/cls

ADD go.mod .
RUN go mod download

RUN go build .

ENTRYPOINT ["./finance"]
