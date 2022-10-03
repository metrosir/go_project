FROM golang:1.16 as builder
WORKDIR /go/src/project
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

#RUN curl https://bootstrap.pypa.io/get-pip.py | python
#RUN pip install supervisor==3.3.5

RUN cd /go/src/project
RUN go mod init
COPY . .
#COPY ./devops/sh/* ./bin/
RUN go build -o ./bin/app ./main.go
#RUN /data/app/bin/app
#RUN /go/src/project/bin/app
CMD ["/go/src/project/bin/app"]