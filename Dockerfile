FROM golang:latest

WORKDIR /NEW

COPY ./ /NEW

RUN go mod download 

ENTRYPOINT go run cmd/app/main.go