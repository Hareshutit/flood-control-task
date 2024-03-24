FROM golang:1.21.1-alpine AS build_stage
COPY . /go/src/vk
WORKDIR /go/src/vk/cmd
RUN go run main.go

