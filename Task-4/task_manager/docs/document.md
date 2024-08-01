# Task Management REST API

## Overview

This project is a simple Task Management REST API implemented using the Go programming language and the Gin framework. It supports basic CRUD operations for managing tasks, including creating, retrieving, updating, and deleting tasks.

## Table of Contents

- [Features](#features)
- [Folder Structure](#folder-structure)
- [Setup and Installation](#setup-and-installation)
- [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
- [Testing](#testing)
- [Technologies Used](#technologies-used)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Create Task**: Add a new task with a title, description, priority, and status.
- **View Tasks**: Retrieve a list of all tasks.
- **View Task by ID**: Retrieve details of a specific task by its ID.
- **Update Task**: Update the details of an existing task.
- **Delete Task**: Remove a task by its ID.

## Folder Structure


My apologies for the confusion. Here's the README.md content formatted as markdown text:

markdown
Copy code
# Task Management REST API

## Overview

This project is a simple Task Management REST API implemented using the Go programming language and the Gin framework. It supports basic CRUD operations for managing tasks, including creating, retrieving, updating, and deleting tasks.

## Table of Contents

- [Features](#features)
- [Folder Structure](#folder-structure)
- [Setup and Installation](#setup-and-installation)


## Features

- **Create Task**: Add a new task with a title, description, priority, and status.
- **View Tasks**: Retrieve a list of all tasks.
- **View Task by ID**: Retrieve details of a specific task by its ID.
- **Update Task**: Update the details of an existing task.
- **Delete Task**: Remove a task by its ID.

## Folder Structure

task_manager/
├── main.go
├── controllers/
│ └── task_controller.go
├── models/
│ └── task.go
├── data/
│ └── task_service.go
├── router/
│ └── router.go
├── docs/
│ └── api_documentation.md
└── go.mod


### File Descriptions

- **main.go**: The main entry point of the application that initializes the server and routes.
- **controllers/task_controller.go**: Contains the HTTP request handlers for each endpoint, orchestrating the data flow between the client and the data layer.
- **models/task.go**: Defines the `Task` struct, representing the structure of a task entity in the application.
- **data/task_service.go**: Implements the business logic and manages the in-memory data store, providing functions for CRUD operations on tasks.
- **router/router.go**: Sets up the Gin router and defines the routing configuration for the API.
- **docs/api_documentation.md**: Contains detailed documentation of the API endpoints, including request and response formats.
- **go.mod**: The Go module file which defines the module's dependencies.

## Setup and Installation

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or higher)
- [Gin Framework](https://github.com/gin-gonic/gin)

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/dipherent1/learning-phase/tree/main/Task-4/task_manager
   cd task_manager```
for more info visit https://documenter.getpostman.com/view/37344769/2sA3kdBdGC

   