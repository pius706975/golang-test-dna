# Golang Test - Digital Nayaka Abhinaya

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/pius706975/golang-test-dna.git
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Running the Application

To start the application, run:

```bash
go run . serve
```

## API Documentation

API documentation is generated using Swagger **v1.16.4**. You can access the documentation by running the server and visiting `<your base url>/docs/index.html` in your browser.

### Generating Swagger Docs

To update Swagger documentation, run:
```bash
swag init
```
Make sure you have installed swaggo globally on your computer.
Read the swaggo documentation [here](https://pkg.go.dev/github.com/swaggo/swag/v2#readme-getting-started)

## Folder Structure

Here's a breakdown of the project folder structure:

- **api/**: Handles route definitions and server setup
  - **routes/**: Defines all API routes from modules
  - **server.go**: Configures and initializes the server

- **cmd/**: Contains command line scripts
  - **command.line.go**: CLI execution entry point

- **config/**: Configuration files, including environment variables
  - **env.go**: Loads and parses environment variables

- **docs/**: API documentation files (Swagger)
  - **swagger.json** and **swagger.yaml**: Swagger specification files

- **middlewares/**: Middleware functions for request handling
  - **middleware.go**: Validate allowed method

- **modules/**: Core application modules
  - **helper.go**
  - **repo.go**
  - **service.go**
  - **controller.go**
  - **routes.go**

- **package/**: Reusable packages
  - **db_struct/**: 
    - **struct.go**
    - **data.go**
  - **utils/**: 
    - **logger.go**

- **main.go**: Application entry point

I also save the fixed dummy data to test the creating data API in **module** directory named **fixed.post.data.json**. Use this json to test the POST request.

## 👨‍💻 Contributor

- Pius Restiantoro - [GitHub](https://github.com/pius706975)