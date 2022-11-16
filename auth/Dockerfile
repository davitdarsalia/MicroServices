FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download


COPY internal ./
COPY pkg ./

RUN go build -o auth-service

EXPOSE 8082

CMD [ "/auth-service" ]