FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod init github.com/c2251730/thamco-sys

RUN go mod tidy 


RUN go build -o order-management-service .

EXPOSE 8080

CMD ["./order-management-service"]
