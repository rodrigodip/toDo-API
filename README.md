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
``

## Testing the Application
Após rodar o projeto, visite [http://localhost:8080/swagger/index.html#](http://localhost:8080/swagger/index.html#) para testar todos os end points.
Também é possível usar ferramentas como [curl](https://curl.se/) ou [postman](https://www.postman.com/) para testar os endpoints. A seguir temos alguns exemplos de comandos da ferramenta `curl`:

**Lembre-se de ajustar os comandos de acordo com sua necessidade.**

- **Criar uma tarefa:**

  ```
  curl -X POST -H "Content-Type: application/json" -d '{"title": "Lavar o carro", "description":"Comprar pasta de polir"}' http://localhost:8080/createTask
  ```

- **Requisitar uma tarefa por ID:**

  ```
    curl -X GET http://localhost:8080/taskById/{id}
  ```

- **Requisitar todas as tarefa:**

  ```
    curl -X GET http://localhost:8080/allTasks
  ```

- **Atualizar uma tarefa:**

  ```
  curl -X PUT -H "Content-Type: application/json" -d '{"title": "Lavar o carro e a kombi", "description":"Comprar pasta de polir e sabão"}' http://localhost:8080/updateTask/{id}
  ```

- **Atualizar uma tarefa como concluída:**

  ```
    curl -X PUT http://localhost:8080/setTaskDone/{id}
  ```

- **Deletar uma tarefa:**

  ```
    curl -X DELETE http://localhost:8080/deleteTask/{id}
  ```


## Data Models

### Requisições

- `title` (string, required): Títulos precisam conter o mínimo de 3 e máximo 30 caracteres.
- `description` (string):  Descrições são **opcionais** e precisam conter no máximo 50 caracteres.

### Respostas

- `id` (integer): Identificador numérico único da tarefa.
- `title` (string, required): Títulos precisam conter o mínimo de 3 e máximo 30 caracteres.
- `description` (string): Descrições são **opcionais** e precisam conter no máximo 50 caracteres.
- `completed` (bool): Identificador lógico de tarefa completada.

