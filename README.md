
# Gotcha - URL Shortener

**Gotcha** é um projeto desenvolvido em **Java** com **Spring Boot** que oferece um serviço de encurtamento de URLs. Ele utiliza **MongoDB** como banco de dados para armazenar as URLs encurtadas e oferece uma API simples para encurtar e redirecionar URLs.

## Funcionalidades

1. **Encurtar URL**: Recebe uma URL completa e retorna uma versão encurtada.
2. **Redirecionar URL**: Redireciona para a URL completa original quando acessada por meio do ID encurtado.

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
    "shortenedUrl": "https://gotcha.com/abc123"
  }
  ```

### 2. Redirecionar URL
- **Endpoint**: `/{id}`
- **Método**: `GET`
- **Descrição**: Redireciona para a URL original baseada no ID encurtado.
- **Exemplo**:
  - Acessando `https://gotcha.com/abc123` redirecionará para `https://exemplo.com`.

## Requisitos

- **Java 17+**
- **Spring Boot 3.x**
- **MongoDB**

## Como rodar o projeto

1. Clone o repositório:
   ```bash
   git clone https://github.com/gabrielalmir/gotcha.git
   ```
2. Entre no diretório do projeto:
   ```bash
   cd gotcha
   ```
3. Configure o acesso ao MongoDB no arquivo `application.properties` ou `application.yml`.
4. Rode o projeto:
   ```bash
   ./mvnw spring-boot:run
   ```

## Estrutura do Projeto

- **Controller**: A camada que expõe os endpoints REST.
  - `UrlController`: Gerencia as operações de encurtar e redirecionar URLs.
  
- **Service**: Lógica de negócios para manipulação de URLs.
  - `UrlService`: Implementa o encurtamento e recuperação de URLs.

- **DTOs**: Objetos de transferência de dados.
  - `ShortenUrlRequest`: Define a estrutura da requisição para encurtar a URL.
  - `ShortenUrlResponse`: Define a estrutura da resposta ao encurtar uma URL.

