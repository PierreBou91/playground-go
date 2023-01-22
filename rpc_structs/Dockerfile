FROM golang:1.19 AS builder

COPY . /usr/local/go/src/rpc_structs/

WORKDIR /usr/local/go/src/rpc_structs/

RUN CGO_ENABLED=0 GOOS=linux go build structs.go

RUN go install rpc_structs
