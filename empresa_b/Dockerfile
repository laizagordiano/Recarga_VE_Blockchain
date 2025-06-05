FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o api-restB ./empresa_b/api-restB.go

EXPOSE 8082
CMD ["./api-restB"]
