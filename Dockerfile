FROM golang:1.23.3-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o auth ./cmd/main.go

EXPOSE 8080

# Adicione permissão explícita para executar o binário
RUN chmod +x /app/auth

CMD ["/app/auth"]
