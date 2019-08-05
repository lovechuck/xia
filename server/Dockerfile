FROM golang:1.12.7-alpine

RUN mkdir -p /data/app

WORKDIR /data/app

COPY . /data/app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

EXPOSE 8080

ENTRYPOINT  ["./main"]