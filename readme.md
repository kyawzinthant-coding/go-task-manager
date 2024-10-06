# Task Manager API with Go, Fiber, GORM, and SQLite

This is a simple RESTful API built with Go using the [Fiber](https://gofiber.io/) web framework, [GORM](https://gorm.io/) ORM for database interactions, and SQLite as the database. This API allows you to perform basic CRUD (Create, Read, Update, Delete) operations on tasks.

## Features

- Create a new task
- Retrieve all tasks
- Retrieve a single task by ID
- Update a task by ID
- Delete a task by ID

## Prerequisites

To run this project, you will need:

- Go (version 1.20+)
- SQLite (installed on your system)
- Go modules (enabled by default in Go 1.11+)

## Project Structure

```bash
.
├── go.mod
├── go.sum
├── main.go           # Entry point for the application
├── models
│   └── task.go       # Task model and database connection logic
├── routes
│   └── task_routes.go # Routes for handling CRUD operations
├── tasks.db          # SQLite database (generated =
└── README.md         
