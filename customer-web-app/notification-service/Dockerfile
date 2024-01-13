FROM golang:latest

WORKDIR /app

COPY ./customer-web-app/notification-service/ .

RUN go mod init github.com/c2251730/thamco-sys

RUN go mod tidy 

RUN go build -o email-microservice .

EXPOSE 80

ENV SMTP_SERVER=smtp.example.com
ENV SMTP_PORT=587
ENV SMTP_USERNAME=your-smtp-username
ENV SMTP_PASSWORD=your-smtp-password

CMD ["./notification-microservice"]
