
# Estágio de build
FROM golang:1.23.1-alpine AS builder

# Definindo o diretório de trabalho
WORKDIR /app

# Copiando os arquivos de dependências
COPY go.mod go.sum ./

# Baixando as dependências
RUN go mod download

# Copiando o código fonte
COPY . .

# Compilando a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o /spl-auth ./cmd/server/main.go

# Estágio final
FROM alpine:latest

# Instalando certificados CA
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiando o binário compilado do estágio de build
COPY --from=builder /spl-auth .
COPY --from=builder /app/.env .

# Expondo a porta
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./spl-auth"]
