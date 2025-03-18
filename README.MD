SPL-Auth Service

Um serviço de autenticação desenvolvido em Go usando o framework Gin.

## Descrição

SPL-Auth é um microserviço de autenticação que fornece funcionalidades de gerenciamento de usuários e autenticação segura.

## Requisitos

- Go 1.23.1 ou superior
- Docker (opcional)
- PostgreSQL (para armazenamento de dados)

## Dependências principais

- Gin Web Framework - Framework web rápido e leve
- GORM - ORM para Go
- JWT-Go - Implementação de JSON Web Tokens
- Validator - Validação de estruturas
- Crypto - Funções criptográficas
- Godotenv - Carregamento de variáveis de ambiente

## Instalação

1. Clone o repositório:
```bash
git clone https://github.com/LissaiDev/spl-auth.git
```

2. Entre no diretório do projeto:
```bash
cd spl-auth
```

3. Instale as dependências:
```bash
go mod download
```

4. Configure as variáveis de ambiente:
```bash
cp .env.example .env
```

## Executando o projeto

### Localmente

1. Certifique-se que o PostgreSQL está rodando
2. Configure as variáveis de ambiente no arquivo .env
3. Execute o servidor:

```bash
go run cmd/server/main.go
```

### Com Docker

```bash
docker-compose up
```

## Endpoints da API

- `POST /auth` - Registro de novo usuário
- `POST /auth/login` - Login de usuário
- `GET /auth` - Obter dados do usuário

## Estrutura do Projeto

```
spl-auth/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   └── services/
├── pkg/
├── docker-compose.yml
└── README.md
```

## Desenvolvimento

Para contribuir com o projeto:

1. Faça um fork do repositório
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Faça commit das suas alterações (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Crie um Pull Request

## Licença

[MIT](https://choosealicense.com/licenses/mit/)

## Contato

Lissai Dev - [@LissaiDev](https://github.com/LissaiDev)

Link do projeto: [https://github.com/LissaiDev/spl-auth](https://github.com/LissaiDev/spl-auth)
