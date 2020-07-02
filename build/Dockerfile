FROM golang:alpine AS build_base

ENV GO111MODULE=on

RUN apk add --no-cache git bash gcc g++ libc-dev && \
    mkdir /calculator-microservice

WORKDIR /calculator-microservice

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o server /calculator-microservice/cmd/server/main.go

EXPOSE 9092

CMD [ "./server" ]
