FROM ubuntu:16.04

ENV GOPATH /go
ENV GOROOT /gotool
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH

COPY ./ $GOPATH/src/github.com/voltento/walletManager/
RUN apt-get update && apt-get install -y git wget
RUN cd /tmp && wget https://dl.google.com/go/go1.11.linux-amd64.tar.gz && tar -xvf go1.11.linux-amd64.tar.gz && mv go /gotool
RUN go get github.com/voltento/walletManager
RUN ln $GOPATH/bin/walletManager /app
COPY .deploy/walletmgr/files /

CMD ["/entrypoint.sh"]
