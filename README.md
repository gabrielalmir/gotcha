# Gotcha - URL Shortener

**Gotcha** é um projeto desenvolvido em **Node.js** com **Express** que oferece um serviço de encurtamento de URLs. Ele utiliza **MongoDB** como banco de dados para armazenar as URLs encurtadas e oferece uma API simples para encurtar e redirecionar URLs.

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
    "shortenedUrl": "http://localhost:3000/abc123"
  }
  ```

### 2. Redirecionar URL
- **Endpoint**: `/{id}`
- **Método**: `GET`
- **Descrição**: Redireciona para a URL original baseada no ID encurtado.
- **Exemplo**:
  - Acessando `http://localhost:3000/abc123` redirecionará para `https://exemplo.com`.

## Requisitos

- **Node.js 14+**
- **Express**
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

3. Instale as dependências:
   ```bash
   npm install
   ```

4. Configure o acesso ao MongoDB:
   - Atualize o arquivo `src/database.ts` com a string de conexão correta para o seu MongoDB.
   ```typescript
   const database = new Database('mongodb://localhost:27017/mydatabase'); // Ajuste para o seu ambiente
   ```

5. Inicie o servidor:
   ```bash
   npm start
   ```

   O servidor será iniciado em `http://localhost:3000`.

## Estrutura do Projeto

- **Controller**: A camada que expõe os endpoints REST.
  - `UrlController`: Gerencia as operações de encurtar e redirecionar URLs.

- **Service**: Lógica de negócios para manipulação de URLs.
  - `UrlService`: Implementa o encurtamento e recuperação de URLs.

- **Model**: Definição da entidade URL no MongoDB.
  - `UrlEntity`: Representa a URL encurtada e sua data de expiração no banco de dados.

- **Utilitários**:
  - `Snowflake`: Gerador de IDs únicos baseado no algoritmo Snowflake.
  - `Base62Converter`: Converte o ID único gerado para uma string Base62 usada nas URLs encurtadas.
