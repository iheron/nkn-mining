FROM golang:1.12-stretch
LABEL maintainer="gdmmx@nkn.org"

ADD ./dist/linux-arm64 /go/src/github.com/nknorg/nkn-mining
WORKDIR /go/src/github.com/nknorg/nkn-mining

EXPOSE 8181
EXPOSE 30001
EXPOSE 30002
EXPOSE 30003
CMD ["./NKNMining", "--remote"]
