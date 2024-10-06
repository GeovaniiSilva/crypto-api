# Etapa de construção
FROM golang:1.19-alpine AS builder

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar os arquivos da aplicação para o contêiner
COPY . .

# Instalar dependências necessárias
RUN apk add --no-cache git

# Compilar a aplicação
RUN go mod tidy
RUN go build -o crypto-api

# Etapa final
FROM alpine:latest

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar o binário compilado da etapa de construção
COPY --from=builder /app/crypto-api .

# Expor a porta 8080
EXPOSE 8080

# Definir o comando de entrada para rodar a aplicação
CMD ["./crypto-api"]
