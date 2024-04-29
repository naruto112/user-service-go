# User Service API

This repository contains the implementation of a User Service API. The API provides endpoints for managing user data.

## Getting Started

To get started with the User Service API, follow these steps:

1. Clone the repository: `git clone https://github.com/your-username/user-service.git`
2. Install the dependencies: `go mod download`
3. Build the binary: `go build`
4. Run the API: `./user-service`

## Endpoints

The User Service API provides the following endpoints:

### Get User

- **Endpoint**: `/users/{id}`
- **Method**: `GET`
- **Description**: Get a user by ID.
- **Response**:
  - `200 OK`: Returns the user data.
  - `400 Bad Request`: If the request is invalid.
  - `404 Not Found`: If the user is not found.

### Get All Users

- **Endpoint**: `/users`
- **Method**: `GET`
- **Description**: Get all users.
- **Response**:
  - `200 OK`: Returns an array of user data.
  - `404 Not Found`: If no users are found.

### Create User

- **Endpoint**: `/users`
- **Method**: `POST`
- **Description**: Create a new user.
- **Request Body**:
  - `name` (string): The name of the user.
  - `email` (string): The email of the user.
- **Response**:
  - `200 OK`: Returns a success message.
  - `400 Bad Request`: If the request is invalid.

### Update User

- **Endpoint**: `/users/{id}`
- **Method**: `PUT`
- **Description**: Update a user by ID.
- **Request Body**:
  - `name` (string): The name of the user.
  - `email` (string): The email of the user.
- **Response**:
  - `200 OK`: Returns a success message.
  - `400 Bad Request`: If the request is invalid.
  - `404 Not Found`: If the user is not found.

### Delete User

- **Endpoint**: `/users/{id}`
- **Method**: `DELETE`
- **Description**: Delete a user by ID.
- **Response**:
  - `200 OK`: Returns a success message.
  - `400 Bad Request`: If the request is invalid.
  - `404 Not Found`: If the user is not found.

## Healthcheck

- **Endpoint**: `/healthcheck`