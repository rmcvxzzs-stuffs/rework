FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY src/ ./src/
COPY config.json ./config.json

WORKDIR /app/src

RUN go build -o /app/lbpk

WORKDIR /app

EXPOSE 8080

CMD ["./lbpk"]
