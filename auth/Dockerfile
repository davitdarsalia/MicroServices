FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Custom steps - Because we have pkg and internal, hey must be copied into workdir app
COPY internal ./
COPY pkg ./

RUN go build -o app .

EXPOSE 8080

CMD ["./app"]
