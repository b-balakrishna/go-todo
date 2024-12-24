# README for Todo List Application

## Overview
This is a full-stack Todo List application built with React for the frontend and Go for the backend. The application allows users to create, view, and delete todo items without the need for a database.

## Project Structure
The project is organized into two main directories: `frontend` and `backend`.

```
todo-list-app
├── backend
│   ├── main.go
│   ├── handlers
│   │   └── todo.go
│   ├── models
│   │   └── todo.go
│   └── README.md
├── frontend
│   ├── src
│   │   ├── App.jsx
│   │   ├── components
│   │   │   └── TodoItem.jsx
│   │   ├── pages
│   │   │   └── TodoList.jsx
│   │   └── main.jsx
│   ├── public
│   │   └── index.html
│   ├── package.json
│   ├── vite.config.js
│   └── README.md
└── README.md
```

## Getting Started

### Prerequisites
- Go (1.16 or later)
- Node.js (14 or later)

### Backend Setup
1. Navigate to the `backend` directory:
   ```
   cd backend
   ```
2. Run the Go server:
   ```
   go run main.go
   ```
3. The server will start on `http://localhost:8080`.

### Frontend Setup
1. Navigate to the `frontend` directory:
   ```
   cd frontend
   ```
2. Install the dependencies:
   ```
   npm install
   ```
3. Start the Vite development server:
   ```
   npm run dev
   ```
4. The application will be available at `http://localhost:3000`.

## API Endpoints
- `GET /todos`: Retrieve the list of todos.
- `POST /todos`: Create a new todo item.
- `DELETE /todos/:id`: Delete a todo item by ID.

## Contributing
Feel free to fork the repository and submit pull requests for any improvements or features.

## License
This project is open-source and available under the MIT License.