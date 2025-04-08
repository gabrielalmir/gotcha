# Gotcha - URL Shortener

**Gotcha** é um projeto desenvolvido em **Go** que oferece um serviço de encurtamento de URLs. Ele utiliza **MongoDB** como banco de dados para armazenar as URLs encurtadas e oferece uma API simples para encurtar e redirecionar URLs.

## Funcionalidades

1. **Encurtar URL**: Recebe uma URL completa e retorna uma versão encurtada.
2. **Redirecionar URL**: Redireciona para a URL completa original quando acessada por meio do ID encurtado.
3. **IDs Únicos**: Utiliza o algoritmo Snowflake para gerar IDs únicos e distribuídos.

## Endpoints

### 1. Encurtar URL
- **Endpoint**: `/shorten-url`
- **Método**: `POST`
- **Descrição**: Encurta uma URL fornecida no corpo da requisição.
- **Exemplo de Requisição**:
  ```json
  {
    "url": "https://exemplo.com"
  }
  ```
- **Exemplo de Resposta**:
  ```json
  {
    "original": "https://exemplo.com",
    "short": "2bX7Pk",
    "created_at": "2024-04-08T21:30:00Z"
  }
  ```

### 2. Redirecionar URL
- **Endpoint**: `/{short}`
- **Método**: `GET`
- **Descrição**: Redireciona para a URL original baseada no ID encurtado.
- **Exemplo**:
  - Acessando `http://localhost:3000/2bX7Pk` redirecionará para `https://exemplo.com`.

## Requisitos

- **Go 1.21+**
- **MongoDB**
- **Docker** (opcional)

## Como rodar o projeto

1. Clone o repositório:
   ```bash
   git clone https://github.com/gabrielalmir/gotcha.git
   ```

2. Entre no diretório do projeto:
   ```bash
   cd gotcha
   ```

3. Instale as dependências:
   ```bash
   go mod download
   ```

4. Configure o MongoDB:
   - Use o Docker Compose fornecido ou configure seu próprio MongoDB
   ```bash
   docker-compose up -d
   ```

5. Configure as variáveis de ambiente:
   ```bash
   cp .env.example .env
   ```

6. Inicie o servidor:
   ```bash
   go run src/main.go
   ```

   O servidor será iniciado em `http://localhost:3000`.

## Estrutura do Projeto

- **Controller**: A camada que expõe os endpoints REST.
  - `URLController`: Gerencia as operações de encurtar e redirecionar URLs.

- **Service**: Lógica de negócios para manipulação de URLs.
  - `URLService`: Implementa o encurtamento e recuperação de URLs.

- **Model**: Definição da entidade URL no MongoDB.
  - `URL`: Representa a URL encurtada no banco de dados.

- **Utilitários**:
  - `Snowflake`: Gerador de IDs únicos baseado no algoritmo Snowflake do Twitter.
  - `Base62`: Converte IDs Snowflake para strings curtas e legíveis.

## Tecnologias Utilizadas

- **Go**: Linguagem de programação
- **Gin**: Framework web
- **MongoDB**: Banco de dados
- **Docker**: Containerização (opcional)
- **Snowflake**: Algoritmo para geração de IDs únicos
