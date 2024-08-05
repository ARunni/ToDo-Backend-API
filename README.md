# TODO APP BACKEND REST API
The To-Do App Backend REST API is designed to manage tasks and to-do items, providing a robust and scalable solution for creating, reading, updating, and deleting (CRUD) tasks.The Todo App developed using Golang, leveraging microservice architecture to ensure scalability, maintainability, and flexibility. This backend service manages tasks, users, and authentication, providing a robust foundation for a Todo application.

## Key Features
1. User Authentication and Authorization:
- **User registration and login.**
- **Secure password hashing.**
- **JWT-based authentication to secure endpoints.**
2. Task Management:
- **Create new tasks with details like title, description, due date.**
- **Retrieve tasks, supporting pagination and filtering.**
- **Update task details.**
- **Delete tasks.**

## Technologies and Libraries
- **Golang**: The primary language used for building the API, chosen for its performance and efficiency.
- **Gin**: A web framework for building RESTful APIs in Golang, providing routing, middleware support, and more.
- **Database**: Postgresql for storing user and task data
- **GORM**: An ORM library for Golang, used for interacting with the database.
- **JWT**: For secure user authentication.
- **Swagger**: For API documentation, making it easier for developers to understand and use the API.
- **Containerization**: Docker for containerizing microservices


## Getting Started

To run the project locally, you can follow these steps:

1. Clone the repository.
2. Set up your environment with the required dependencies, including Golang, PostgreSQL and Docker.
3. Configure your environment variables (e.g., database credentials).
4. Build and run the project.


# Environment Variables

Before running the project, you need to set the following environment variables with your corresponding values:

## Api-service
- `PORT`   : Base Url
- `TODO_SVC_URL` : localhost:port

## ToDo-service

- `PORT`   : Base Url
- `DB_HOST` : localhost
- `DB_NAME` : Database name
- `DB_USER` : Database user
- `DB_PORT` : Database port
- `DB_PASSWORD` : Database password

Make sure to provide the appropriate values for these environment variables to configure the project correctly.

