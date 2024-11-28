# API-CRUD-GO

## 📋 Descrição do Projeto
Este é um projeto de API RESTful desenvolvido em Golang, utilizando o framework **Gin-Gonic**. Foi projetado para atender às requisições do front-end e permitir a comunicação eficiente com o banco de dados. A API implementa funcionalidades de CRUD (Criar, Ler, Atualizar, Deletar) para gerenciar clientes e produtos.

A estrutura do projeto foi organizada para facilitar a manutenção, escalabilidade e entendimento do código.

## 🛠️ Tecnologias Utilizadas
- **Go (Golang)**: Linguagem de programação utilizada no back-end.
- **Gin**: Framework web para a construção da API RESTful.
- **Postman**: Ferramenta para testes e validação dos endpoints.
- **MySQL Driver** ([github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)): Driver para conectar o Golang ao banco de dados MySQL.
- **Validator** ([github.com/go-playground/validator/v10](https://github.com/go-playground/validator/v10)): Biblioteca para validação de dados de entrada.
- **Gin-CORS**: Middleware para suporte a CORS (Cross-Origin Resource Sharing).

## 📁 Estrutura do Projeto
A estrutura do projeto é a seguinte:
API-CRUD-GO/
├── controllers/
│   ├── cliente.go         # Controlador para operações de clientes
│   ├── produtos.go        # Controlador para operações de produtos
├── database/
│   ├── connectionBanco.go # Configuração e conexão com o banco de dados
├── modelos/
│   ├── modelos.go         # Definições de modelos (Clientes, Produtos)
├── rotas/
│   ├── rotas.go           # Definição de rotas da API
├── go.mod                 # Arquivo de dependências do Go
├── go.sum                 # Resumo das dependências
├── main.go                # Arquivo principal da aplicação
├── README.md              # Documentação do projeto

## 🚀 Endpoints da API
Abaixo estão os principais endpoints disponíveis:

### **Clientes**
| Método | Endpoint         | Descrição                    |
|--------|------------------|------------------------------|
| GET    | /clientes         | Lista todos os clientes      |
| GET    | /clientes/:id     | Consulta um cliente específico |
| POST   | /clientes         | Cria um novo cliente         |
| PUT    | /clientes/:id     | Atualiza um cliente específico |
| DELETE | /clientes/:id     | Remove um cliente            |

### **Produtos**
| Método | Endpoint         | Descrição                    |
|--------|------------------|------------------------------|
| GET    | /produtos         | Lista todos os produtos      |
| GET    | /produtos/:id     | Consulta um produto específico |
| POST   | /produtos         | Cria um novo produto         |
| PUT    | /produtos/:id     | Atualiza um produto específico |
| DELETE | /produtos/:id     | Remove um produto            |

## ⚙️ Configuração e Execução

### **Pré-requisitos**
Certifique-se de ter os seguintes softwares instalados:
- **Go**
- **Postman** (opcional, para testes)

### **Passo a passo**

**1. Clone o repositório:**
   git clone https://github.com/c0lllt/api-crud-go



**cd API-CRUD-GO**
**Configure as dependências:**


go mod tidy
**Execute o servidor:**


go run main.go
**A API estará disponível em: http://localhost:8081**

## 🧪 Testes ##
Os endpoints podem ser testados usando o Postman ou qualquer outra ferramenta de sua escolha. Certifique-se de configurar corretamente o corpo das requisições para os métodos POST e PUT.

Exemplo de Requisição para Criar um Cliente (POST /clientes)
{
  "nome": "Testando",
  "cpf": "000.000.000-00",
  "senha": "123@456"
}

Exemplo de Requisição para Criar um Produto (POST /produtos)
{
  {
   "nome": "Tokio",
    "descricao":"Atum, Cream Chesse, Alho Poró, Parmesão, Molho Tarê e Gergelin",
    "valor":79.99
  
}
}
## 🗂️ Contribuição ##
Contribuições são bem-vindas! 


## 👤 Autor
Desenvolvido por **Paulo Henrique Nunes**.  
Entre em contato via [LinkedIn](www.linkedin.com/in/paulo-henrique-nunes-55b5ab22b) ou [E-mail](paulo.hnrnunes@gmail.com).

