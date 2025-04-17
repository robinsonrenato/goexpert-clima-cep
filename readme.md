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

---

## 🏗️ Estrutura do Projeto

```plaintext
.
├── cmd
│   └── server
│       └── main.go          # Ponto de entrada da aplicação
├── internal
│   ├── controller
│   │   └── weather_handler.go # Handler HTTP para o endpoint de clima
│   ├── service
│   │   └── weather_service.go # Lógica de negócio para consulta de clima
│   ├── model
│   │   └── client.go         # Interfaces para os clientes externos
│   └── util
│       └── temperature.go    # Funções utilitárias para conversão de temperatura
├── test
│   ├── controller
│   │   └── weather_handler_test.go # Testes para o handler HTTP
│   ├── service
│   │   └── weather_service_test.go # Testes para o serviço de clima
│   └── util
│       └── util_test.go            # Testes para funções utilitárias
├── Dockerfile                     # Configuração para construir a imagem Docker
├── docker-compose.yml             # Configuração para gerenciar os serviços com Docker Compose
└── go.mod                         # Gerenciamento de dependências do Go
```

---

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
