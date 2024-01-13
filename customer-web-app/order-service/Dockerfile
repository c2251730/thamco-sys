FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod init github.com/c2251730/thamco-sys

RUN go mod tidy 

RUN go build -o order-microservice .

EXPOSE 8082

CMD ["./order-microservice"]
