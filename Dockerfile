FROM golang:1.24.3-alpine3.21

# Airのインストール
RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/src

CMD [ "air" ]
