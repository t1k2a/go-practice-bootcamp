# Goのビルドステージ
FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o server .

# 実行ステージ（小さいイメージにコピー）
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/server .
COPY .env .

# ポートを明示（任意）
EXPOSE 8080

CMD ["./server"]