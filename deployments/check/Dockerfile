# 1 шаг - сборки
FROM golang:1.21.1-alpine AS build_stage
COPY . /go/src/vk
WORKDIR /go/src/vk/cmd
RUN go env -w GO111MODULE=auto
RUN go install .

# 2 шаг
FROM alpine AS run_stage
WORKDIR /app_binary
COPY --from=build_stage /go/bin/ /app_binary/
RUN chmod +x ./
ENTRYPOINT ./