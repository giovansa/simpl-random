FROM golang:1.23.0-alpine3.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8081

CMD ["./main"]
