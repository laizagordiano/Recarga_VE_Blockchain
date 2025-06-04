FROM golang:1.21-alpine

WORKDIR /app

# Copia os arquivos de dependência primeiro para cache do Docker
COPY go.mod go.sum ./

# Baixa as dependências
RUN go mod tidy

# Copia o restante do código
COPY . .

# Build do binário da aplicação
RUN go build -o app .

# Expõe a porta usada pela aplicação
EXPOSE 8083

# Comando para rodar a aplicação
CMD ["./app"]
