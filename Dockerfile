FROM golang:1.24.3 AS builder

WORKDIR /app

# go.mod / go.sum だけ先にコピーしてキャッシュ効かせる
COPY go.mod go.sum ./
RUN go mod download

# 残りのソースをコピー
COPY src/ ./src

# ビルド実行（main.go は src/main.go にある）
WORKDIR /app/src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main .

# 実行用イメージ
FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
