FROM golang:latest

WORKDIR /app

COPY ./staff-web-app/customer-management-service/ .

RUN go mod init github.com/c2251730/thamco-sys

RUN go mod tidy 

RUN go build -o customer-management-service .

EXPOSE 8080

CMD ["./customer-management-service"]
