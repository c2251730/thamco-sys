FROM golang:latest

WORKDIR /app

COPY /main.go .

RUN go mod init github.com/c2251730/thamco-sys

RUN go mod tidy 

RUN go build main.go

EXPOSE 8082

CMD ["./main"]
