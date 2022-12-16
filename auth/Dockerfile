FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY internal ./
COPY pkg ./

RUN go build -o app .

EXPOSE 8080

CMD ["./app"]
