# Meu CRUD em Go

This is a comprehensive guide for the "Meu CRUD em Go" project in Go, an example application that implements the basic CRUD (Create, Read, Update, Delete) operations for users. The project includes a Dockerfile to facilitate running it in containers.

## Information

- **Title**: Meu CRUD em Go | Lucariom
- **Version**: 1.0
- **Host**: localhost:8080

## Prerequisites

Before getting started, make sure you have the following prerequisites installed on your system:

- [Go](https://golang.org/dl/): The Go programming language.
- [Docker](https://www.docker.com/get-started): Docker is required if you wish to run the application in a container.

## Installation

Follow the steps below to install the project in your development environment:

1. **Clone the repository:**

   ```
   git clone https://github.com/LucasAntonioC-137/crud-go.git
   ```

2. **Navigate to the project directory:**

   ```
   cd crud-go
   ```

3. **Build the application using Docker Compose:**

   ```
   docker compose up
   ```

## Running the Application

After installation, you can run the Meu CRUD em Go application with the following command (if you want to run it directly with Golang):

```
docker container run --name Meu CRUDgo -p 27017:27017 -d mongo

go run main.go
```

The application will be accessible at `http://localhost:8080`.

## Testing the Application

If you prefer, after running the project, visit: http://localhost:8080/swagger/index.html# to see and test all the route contracts.

This project has been updated and can now only be used if you have a token. To generate the first token, you need to build the project using the init-mongo.js script. To create the admin login, use the login route with the email admin@crud.com and password admin123# to get the token in the header. All routes needs the token, except the /login.

The Meu CRUD em Go application offers REST endpoints for creating, listing, updating, and deleting users. You can use tools like [curl](https://curl.se/) or [Postman](https://www.postman.com/) to test the endpoints. Here are some `curl` command examples for testing the endpoints:

- **Create a user:**

  ```
  curl -X POST -H "Content-Type: application/json" -d '{"name": "João", "email": "joao@example.com", "age": 30, "password": "password$#@$#323"}' http://localhost:8080/createUser
  ```

- **Update a user:**

  ```
  curl -X PUT -H "Content-Type: application/json" -d '{"name": "João Silva"}' http://localhost:8080/updateUser/{userId}
  ```

- **Delete a user:**

  ```
  curl -X DELETE http://localhost:8080/deleteUser/{userID}
  ```

Remember to adjust the commands according to your needs and requirements.

### Note

- For authentication, you should include the access token in the `Authorization` header as "Bearer <Insert access token here>" for protected endpoints.

The API offers the following endpoints:

1. **POST /createUser**
   - Description: Create a new user with the provided user information.
   - Parameters:
      - `userRequest` (body, required): User information for registration.
   - Responses:
      - 200: OK (User created successfully)
      - 400: Bad Request (Request error)
      - 500: Internal Server Error (Internal server error)

2. **DELETE /deleteUser/{userId}**
   - Description: Delete a user based on the provided ID parameter.
   - Parameters:
      - `userId` (path, required): ID of the user to be deleted.
   - Responses:
      - 200: OK (User deleted successfully)
      - 400: Bad Request (Request error)
      - 500: Internal Server Error (Internal server error)

3. **GET /getUserByEmail/{userEmail}**
   - Description: Retrieve user details based on the email provided as a parameter.
   - Parameters:
      - `userEmail` (path, required): Email of the user to be retrieved.
   - Responses:
      - 200: User information retrieved successfully
      - 400: Error: Invalid user ID
      - 404: User not found

4. **GET /getUserById/{userId}**
   - Description: Retrieve user details based on the user ID provided as a parameter.
   - Parameters:
      - `userId` (path, required): ID of the user to be retrieved.
   - Responses:
      - 200: User information retrieved successfully
      - 400: Error: Invalid user ID
      - 404: User not found

5. **POST /login**
   - Description: Allow a user to log in and receive an authentication token.
   - Parameters:
      - `userLogin` (body, required): User login credentials.
   - Responses:
      - 200: Login successful, authentication token provided
      - 403: Error: Invalid login credentials

6. **PUT /updateUser/{userId}**
   - Description: Update user details based on the ID provided as a parameter.
   - Parameters:
      - `userId` (path, required): ID of the user to be updated.
      - `userRequest` (body, required): User information for update.
   - Responses:
      - 200: OK (User updated successfully)
      - 400: Bad Request (Request error)
      - 500: Internal Server Error (Internal server error)

## License

This project is distributed under the MIT license. Please refer to the LICENSE file for more details.

---

We hope this Swagger documentation has been helpful in understanding and interacting with the API of the Meu CRUD em Go project in Go. If you have any questions or need additional support, please don't hesitate to reach out. Enjoy using the API!
