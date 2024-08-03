
# Task Management API with MongoDB

## Project Overview

This project extends an existing Task Management API by integrating MongoDB for persistent data storage using the Mongo Go Driver. The API provides CRUD operations for managing tasks and ensures data persistence across API restarts.

## Project Structure

```
task_manager/
├── main.go
├── config/
│   └── config.go
├── controllers/
│   └── task_controller.go
├── models/
│   └── task.go
├── services/
│   └── task_service.go
├── router/
│   └── router.go
├── docs/
│   └── api_documentation.md
└── go.mod
```

### Key Components

- **main.go**: Entry point of the application. Initializes the MongoDB connection and starts the Gin server.
- **config/**: Contains configuration logic, including MongoDB connection settings.
- **controllers/**: Handles incoming HTTP requests and invokes appropriate service methods.
- **models/**: Defines the data structures used in the application.
- **services/**: Contains business logic and data manipulation functions. Implements MongoDB operations.
- **router/**: Sets up the Gin router and defines the API routes.
- **docs/**: Contains project documentation.

## Getting Started

### Prerequisites

- Go (version 1.16 or later)
- MongoDB (either locally or using a cloud provider)
- Gin (Go web framework)

### Installation

1. **Clone the Repository**

   ```sh
   git clone <repository-url>
   cd task_manager
   ```

2. **Set Up Environment Variables**

   Create a `.env` file in the root directory with the following content:

   ```
   MONGODB_URI=<your-mongodb-uri>
   ```

3. **Install Dependencies**

   ```sh
   go mod tidy
   ```

4. **Run the Application**

   ```sh
   go run main.go
   ```

## API Endpoints

### List All Tasks

- **URL**: `/tasks`
- **Method**: `GET`
- **Response**: List of all tasks.

### Create a New Task

- **URL**: `/tasks`
- **Method**: `POST`
- **Request Body**:

  ```json
  {
    "title": "Task Title",
    "description": "Task Description",
    "priority": "High",
    "status": "Pending"
  }
  ```

- **Response**: Created task object.

### Get Task by Title

- **URL**: `/tasks/:title`
- **Method**: `GET`
- **Response**: Task object with the specified title.

### Update Task

- **URL**: `/tasks/:title`
- **Method**: `PUT`
- **Request Body**: Fields to update in the task.

- **Response**: Update result object.

### Delete Task

- **URL**: `/tasks/:title`
- **Method**: `DELETE`
- **Response**: Deletion success message.

## Configuration

### MongoDB Connection

The MongoDB connection is configured in the `config/config.go` file. The connection URI is retrieved from the environment variables. Make sure to set the `MONGODB_URI` environment variable before running the application.

## Error Handling

Proper error handling is implemented throughout the application, especially for MongoDB operations. Errors are logged, and appropriate HTTP status codes are returned.

## Testing

To test the API endpoints, use tools like Postman or cURL. Verify that tasks can be created, retrieved, updated, and deleted successfully. You can also query the MongoDB database directly or use MongoDB Compass to ensure data correctness.

