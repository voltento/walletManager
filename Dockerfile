FROM ubuntu:16.04

WORKDIR /go/src/app
COPY . ./project

ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN apt-get update && apt-get install -y golang-go git
RUN go get github.com/voltento/walletManager

CMD ["app"]