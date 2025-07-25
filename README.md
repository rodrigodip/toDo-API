# 🚀 Desafio Técnico - API de Gerenciamento de Tarefas

Este projeto é um desafio técnico com o objetivo de praticar a criação de uma API REST em Go, utilizando um framework web (sugestão: [Gin](https://github.com/gin-gonic/gin)).

---

## 🎯 Objetivo

Desenvolver uma API para gerenciar tarefas (ToDos), com suporte às operações básicas de:

- Criar uma nova tarefa
- Listar todas as tarefas
- Obter uma tarefa por ID
- Atualizar uma tarefa existente
- Remover uma tarefa

---

## 🛠 Requisitos

- Utilize um framework web em Go (preferencialmente *Gin*).
- A persistência dos dados pode ser feita *em memória*.
- Organize o projeto de forma clara e modular.
- Versione o projeto com Git e publique em um repositório no GitHub.
- Inclua este README.md com:
  - Instruções para rodar a aplicação
  - Descrição da API
  - Exemplos de requisição (curl ou Postman)

---

## ▶️ Como rodar o projeto

### Pré-requisitos

- Go 1.20 ou superior

### Passos

1. Clone o repositório
```bash
git clone https://github.com/rodrigodip/toDo-API.git
cd toDo-API
```
2. Instale as dependências
```bash
go mod tidy
```

3. Rode a aplicação
```bash
go run main.go
```
