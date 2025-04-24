# GoExpert Clima CEP - Cloud Run

Este projeto é uma aplicação em Go que consulta informações climáticas com base em um CEP fornecido. Ele utiliza serviços externos para buscar a cidade correspondente ao CEP e a temperatura atual da cidade.

## 🚀Funcionalidades

- Consulta a cidade correspondente a um CEP usando um serviço externo `ViaCEP`.
- Consulta a temperatura atual da cidade usando uma API de clima `weatherapi`.
- Retorna as informações climáticas em graus Celsius, Fahrenheit e Kelvin.

---

## 📋Requisitos

- **Go**: Versão 1.20 ou superior.
- **Docker**: Para executar o projeto em contêineres.
- **Docker Compose**: Para gerenciar múltiplos serviços.

---

## 🏃‍♂️Como compilar e executar localmente

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/seu-usuario/seu-repositorio.git
   go mod tidy
   go run ./cmd/server
   ```

4. **Teste no `Navegador`:**
   ```bash
   http://localhost:8080/weather?cep=01001000
   ```

5. **Teste no do Deploy do Google Cloud:**
   ```bash
   https://weather-api-670673982186.us-central1.run.app/weather?cep=01001000
   ```

---

## 🧪Como executar os testes

1. **Execute todos os testes:**
   ```bash
   go test ./...
   ```

---

## 🏃‍♂️Como executar com Docker

1. **Construa os serviços:**
   ```bash
   docker-compose build
   docker-compose up
   ```

3. **Teste o endpoint:**
   Após iniciar os serviços, o servidor estará disponível em:
   ```bash
   http://localhost:8080/weather?cep=01001000
   ```

4. **Pare os serviços:**
   Quando terminar, pare os serviços com:
   ```bash
   docker-compose down
   ```

---

## 📡Retorno do Endpoint

### Requisição bem-sucedida:
**GET** `/weather?cep=01001000`

**Resposta:**
```json
{
  "city": "São Paulo",
  "temp_c": 25.0,
  "temp_f": 77.0,
  "temp_k": 298.15
}
```

### CEP inválido:
**GET** `/weather?cep=123`

**Resposta:**
```json
{
  "message": "invalid zipcode"
}
```

**Status HTTP:** `422 Unprocessable Entity`

### CEP não encontrado:
**GET** `/weather?cep=99999999`

**Resposta:**
```json
{
  "message": "can not find zipcode"
}
```

**Status HTTP:** `404 Not Found`

### Erro interno:
**GET** `/weather?cep=01001000` (quando ocorre um erro inesperado)

**Resposta:**
```json
{
  "message": "internal error"
}
```

**Status HTTP:** `500 Internal Server Error`


# Deploy no Google Cloud Run

## Pré-requisitos
- Conta no Google Cloud Platform
- Google Cloud SDK instalado
- Projeto criado no Google Cloud
- Faturamento habilitado no projeto

## Passos para Deploy

1. **Instalar o Google Cloud SDK**
   - Baixe o instalador: https://cloud.google.com/sdk/docs/install
   - Execute o instalador
   - Abra um novo terminal após a instalação

2. **Configurar o ambiente**
   ```bash
   gcloud auth login
   gcloud init
   ```

2. **Testando projeto**
   ```bash
   https://weather-api-670673982186.us-central1.run.app/weather?cep=01001000
   ```
   
---
## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
