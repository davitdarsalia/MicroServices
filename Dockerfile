FROM golang:latest

WORKDIR /app


COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

ENV PORT 8080

CMD [ "/bank-api" ]