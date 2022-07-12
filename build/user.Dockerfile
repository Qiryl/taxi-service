FROM golang:1.18.3-alpine3.16

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main ./cmd/user/main.go

EXPOSE 8080

CMD ./main