FROM golang:1.18.3-alpine3.16

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download
RUN go build -o client ./internal/user/cmd/client/client.go
RUN go build -o server ./internal/user/cmd/server/server.go

EXPOSE 8080

CMD ./client & ./server 