# Estágio de construção
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

# Copiar os arquivos go.mod e go.sum
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod download

# Copiar o código restante
COPY . .

# Rodar go mod tidy para limpar dependências desnecessárias
RUN go mod tidy

# Compilar a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o /pbl-recarga ./cmd/server

# Estágio final
FROM alpine:latest

# Instalar certificados necessários
RUN apk --no-cache add ca-certificates

# Copiar o binário gerado para o estágio final
COPY --from=builder /pbl-recarga /pbl-recarga

# Expor a porta da aplicação
EXPOSE 8080

# Definir o comando para iniciar a aplicação
CMD ["/pbl-recarga"]
