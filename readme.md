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

### Healthcheck

- **Endpoint**: `/health-check`
- **Method**: `GET`
- **Description**: Check the health of the API.
- **Response**:
  - `200 OK`: Returns a success message.

### Login

- **Endpoint**: `/users/login`
- **Method**: `POST`
- **Description**: Authenticate a user and generate a token.
- **Request Body**:
  - `email` (string): The email of the user.
  - `password` (string): The password of the user.
- **Response**:
  - `200 OK`: Returns a token for authentication.
  - `400 Bad Request`: If the request is invalid.

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
  - `name` (string): The updated name of the user.
  - `email` (string): The updated email of the user.
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

## License

This project is licensed under the [MIT License](LICENSE).

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## Contact

For any inquiries or questions, please contact [renatorock3@hotmail.com](mailto:renatorock3@hotmail.com).