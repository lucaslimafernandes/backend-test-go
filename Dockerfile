# Etapa de build
FROM golang:1.22 AS builder

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia os arquivos do projeto para o contêiner
COPY main.go .
COPY services/rb/services.go ./services/rb/

# Compila os binários
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/main main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/services ./services/rb/services.go

# Etapa final
FROM alpine:latest

# Copia os binários da etapa de build
COPY --from=builder /bin/main /bin/main
COPY --from=builder /bin/services /bin/services

# Copia o script de entrada
COPY entrypoint.sh /entrypoint.sh

# Dá permissão de execução para o script de entrada
RUN chmod +x /entrypoint.sh

# Define o script de entrada como ponto de entrada do contêiner
ENTRYPOINT ["/entrypoint.sh"]
