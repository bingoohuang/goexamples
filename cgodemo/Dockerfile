FROM ubuntu:18.04

ARG GO=1.15.6

WORKDIR /tmp

RUN apt update && apt upgrade -y && apt install -y wget make gcc vim binutils \
    && wget https://dl.google.com/go/go${GO}.linux-amd64.tar.gz  \
    && tar -xvf go${GO}.linux-amd64.tar.gz \
    && mv go /usr/local \
    && echo export GOROOT=/usr/local/go >> ~/.bashrc  \
    && echo export GOPATH=\$HOME/go >> ~/.bashrc \
    && echo export PATH=\$GOPATH/bin:\$GOROOT/bin:\$PATH >> ~/.bashrc

WORKDIR /src
