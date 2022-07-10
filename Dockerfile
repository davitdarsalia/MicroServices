FROM golang:1.18.3-alpine3.16

ADD . /app
WORKDIR /app

RUN go build -o main .
CMD ["/app/main"]