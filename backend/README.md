# FILE: /todo-list-app/todo-list-app/backend/README.md

# Todo List Application - Backend

This is the backend part of the Todo List application built with Go. It provides a simple API to manage todo items without using a database.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:
   ```
   git clone <repository-url>
   cd todo-list-app/backend
   ```

2. Install dependencies (if any):
   ```
   go mod tidy
   ```

### Running the Server

To run the backend server, use the following command:

```
go run main.go
```

The server will start on `http://localhost:8080`.

### API Endpoints

- `GET /todos` - Retrieve all todo items.
- `POST /todos` - Create a new todo item.
- `DELETE /todos/{id}` - Delete a todo item by ID.

### Example Requests

- **Get Todos**
  ```
  curl -X GET http://localhost:8080/todos
  ```

- **Create Todo**
  ```
  curl -X POST http://localhost:8080/todos -d '{"title": "New Todo", "completed": false}' -H "Content-Type: application/json"
  ```

- **Delete Todo**
  ```
  curl -X DELETE http://localhost:8080/todos/1
  ```

## License

This project is licensed under the MIT License.