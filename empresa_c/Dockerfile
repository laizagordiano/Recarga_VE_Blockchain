FROM golang:1.24-alpine
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o api-restC ./empresa_c/api-restC.go

EXPOSE 8083
CMD ["./api-restC"]
