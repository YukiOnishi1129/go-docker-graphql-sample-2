FROM golang:1.18.0-alpine3.15

RUN apk update && apk add git
RUN mkdir /app
# ワーキングディレクトリの設定
WORKDIR /app

COPY . /app

RUN go mod tidy && \
    go install github.com/cosmtrek/air@v1.27.3