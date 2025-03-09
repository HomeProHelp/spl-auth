# SPL-Auth Service

Um serviço de autenticação desenvolvido em Go usando o framework Gin.

## Descrição

SPL-Auth é um microserviço de autenticação que fornece funcionalidades de gerenciamento de usuários e autenticação segura.

## Requisitos

- Go 1.23.1 ou superior
- Docker (opcional)

## Dependências principais

- Gin Web Framework
- Validator
- Crypto

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

## Executando o projeto

### Localmente

```bash
go run cmd/server/main.go
```

### Com Docker

```bash
docker-compose up
```

## Desenvolvimento

Para contribuir com o projeto:

1. Faça um fork do repositório
2. Crie uma branch para sua feature
3. Faça commit das suas alterações
4. Push para a branch
5. Crie um Pull Request

## Licença

[MIT](https://choosealicense.com/licenses/mit/)
