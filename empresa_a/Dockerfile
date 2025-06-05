FROM golang:1.24-alpine
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o api-rest ./empresa_a/api-rest.go

EXPOSE 8081
CMD ["./api-rest"]
