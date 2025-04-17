# Use uma imagem base do Go
FROM golang:1.23

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie os arquivos do projeto para o contêiner
COPY . .

# Baixe as dependências do projeto
RUN go mod tidy

# Compile o projeto
RUN go build -o server ./cmd/server

# Exponha a porta que o servidor usará
EXPOSE 8080

# Comando para executar o servidor
CMD ["./server"]
