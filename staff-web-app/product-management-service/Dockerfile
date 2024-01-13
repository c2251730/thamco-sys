FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod init github.com/c2251730/thamco-sys

RUN go mod tidy 

RUN go build -o product-management-service .

EXPOSE 8080

CMD ["./product-management-service"]
