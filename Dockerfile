FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 go build -o main ./cmd

FROM debian:latest

RUN apt-get update && \
    apt-get install -y ca-certificates wget && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /root/

COPY --from=builder /app/main .

RUN wget https://dl.google.com/go/go1.22.0.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz && \
    rm go1.22.0.linux-amd64.tar.gz && \
    export PATH=$PATH:/usr/local/go/bin

ENV PATH="/usr/local/go/bin:${PATH}"

EXPOSE 8080

# Команда по умолчанию
CMD ["./main"]
