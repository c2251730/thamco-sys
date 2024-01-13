FROM golang:latest

WORKDIR /app

COPY ./customer-web-app/user-service/ .

RUN go mod init github.com/c2251730/thamco-sys

RUN go mod tidy 

RUN go build -o user-microservice .

EXPOSE 8080

CMD ["./user-microservice"]
