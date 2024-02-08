# Go Communication Service

Este é um serviço em Go que permite enviar mensagens para canais do Telegram e Slack através de uma API HTTP.

## Instalação e Execução

1. Clone o repositório:

```bash
git clone git@github.com:gabicmontes/communications.git
```

2. Acesse o diretório do projeto:

```bash
cd communications
```

3. Crie um arquivo `.env` na raiz do projeto com as variáveis definidas no arquivo `.env.example`.

```bash
cp .env.example .env
```

4. Instale as dependências do projeto:

```bash
go mod tidy
```

5. Execute o projeto:

```bash
go run main.go
```

## Utilização

### Enviar mensagem para o Telegram

- **URL**

  `http://localhost:8080/telegram`

- **Método:**

      `POST`

- **Corpo da requisição:**

 ```json
    {
        "group_id": "ID do grupo",
        "message": "Hello, World!"
    }
 ```

- **Resposta de Sucesso:**

  ```json
    {
        "message": "Mensagem enviada com sucesso!"
    }
  ```

- **Resposta de Erro:**
  ```json
    {
        "error": "Um erro ocorreu ao enviar a mensagem."
    }
    ```

### Enviar mensagem para o Slack

- **URL**

  `http://localhost:8080/slack`

- **Método:**

      `POST`

- **Corpo da requisição:**

  ```json
    {
        "channel": "ID do canal",
        "message": "Hello, World!",
        "title": "Título da mensagem",
        "color": "#36a64f"
    }
  ```
