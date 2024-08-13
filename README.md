# API CRUD em Go

Esta é uma API CRUD (Create, Read, Update, Delete) desenvolvida em Go, utilizando o framework Gin e o MongoDB como banco de dados. A API permite a manipulação de um catálogo de produtos, oferecendo endpoints para adicionar, listar, atualizar e remover produtos.

## Tecnologias

- **Golang:** Linguagem de programação utilizada para desenvolver a API.
- **Gin:** Framework web para gerenciamento das rotas e middlewares.
- **MongoDB:** Banco de dados NoSQL para armazenamento dos dados dos produtos.
- **Postman:** Ferramenta para testar e validar os endpoints da API.

## Funcionalidades

- **Criação de Produtos:** Adicione novos produtos ao banco de dados.
- **Listagem de Produtos:** Liste todos os produtos cadastrados.
- **Atualização de Produtos:** Atualize as informações de um produto existente.
- **Remoção de Produtos:** Delete um produto do banco de dados.

## Estrutura do Projeto

- `handler/`: Contém os manipuladores de rotas para as operações CRUD.
  - `create.go`: Manipula a criação de novos produtos.
  - `delete.go`: Manipula a exclusão de produtos.
  - `list.go`: Manipula a listagem de produtos.
  - `update.go`: Manipula a atualização de produtos.
- `database/`: Configura a conexão com o MongoDB.
  - `database.go`: Gerencia a conexão com o banco de dados.
- `models/`: Define a estrutura do modelo `Produto`.
  - `produto.go`: Estrutura do modelo de produto.
- `main.go`: Arquivo principal que inicia o servidor e configura as rotas.
- `go.mod`: Arquivo de módulo do Go.
- `go.sum`: Resumo das dependências do Go.

## Configuração

1. **Clone o Repositório:**

   ```bash
   git clone https://github.com/JuJurubeba/API.go.git
   cd API.go
Instale as Dependências:


Copiar código:
    
    go mod tidy

Configure o Banco de Dados:

Substitua a URI de conexão com o MongoDB em database/mongo.go pela sua própria URI do MongoDB.
Execute o Servidor:


Copiar código:

    go run main.go


O servidor estará disponível em http://localhost:8080.

Testando a API
Você pode usar o Postman para testar os seguintes endpoints:

POST /produtos: Adiciona um novo produto.
GET /produtos: Lista todos os produtos.
PUT /produtos/
: Atualiza um produto existente.
DELETE /produtos/
: Remove um produto pelo ID.




Contribuições são bem-vindas! Se você encontrar algum problema ou tiver sugestões de melhorias, sinta-se à vontade para abrir uma issue ou enviar um pull request.
