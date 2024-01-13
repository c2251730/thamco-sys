FROM golang:latest

WORKDIR /app

COPY ./customer-web-app/product-service/ .

RUN go mod init github.com/c2251730/thamco-sys

RUN go mod tidy 

RUN go build -o product-microservice .

EXPOSE 8081

CMD ["./product-microservice"]
