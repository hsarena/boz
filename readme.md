# Boz Projects Documentation

## Overview

A REST API project that allows you to manage a herd of sheep. This API provides functionalities for creating sheep entities with attributes such as name, color, wool type, and breed. Additionally, calling the /baa endpoint results in a sheep responding with a cheerful "baa". Dont mind, this is just for fun

## API Endpoints
### 1. Create a Sheep
- **Endpoint:** POST /sheeps
- **Description:** Create a new sheep with attributes such as name, color, wool type, and breed.
- **Request Payload:**
    ```json
    {
        "name": "Fluffy",
        "color": "White",
        "wool": "Carpet",
        "breed": "Merino"
    }
- **Example Response:**
    ```json
    {
        "status": "success",
        "sheep": {
            "id": 1,
            "name": "Fluffy",
            "color": "White",
            "wool": "Carpet",
            "breed": "Merino",
            "created_at": "2023-11-20T12:00:00Z"
        }
    }
### 2. Get All Sheep
- **Endpoint:** GET /sheeps
- **Description:** Retrieve a list of all sheep in the herd.
- **Example Response:**
    ```json
    {
        "status": "success",
        "results": 1,
        "data": [
            {
            "id": 1,
            "name": "Fluffy",
            "color": "White",
            "wool": "Carpet",
            "breed": "Merino",
            "created_at": "2023-11-20T12:00:00Z"
            }
        ]
    }
### 3. Get Sheep by ID
- **Endpoint:** GET /sheeps/:id
- **Description:** Retrieve details of a specific sheep by ID.
- **Example Response:**
    ```json
    {
        "status": "success",
        "sheep": {
            "id": 1,
            "name": "Fluffy",
            "color": "White",
            "wool": "Carpet",
            "breed": "Merino",
            "created_at": "2023-11-20T12:00:00Z"
        }
    }
### 4. Sheep Responds to /baa
- **Endpoint:** POST /baa
- **Description:** The sheep responds with a cheerful "baa."
- **Example Response:**
    ```json
    {
        "status": "success",
        "data": {
            "Sheep": "Fluffy",
            "Baa": "Baaa"
        }
    }

## Environment Variables


####  HOST

- **Description** The host where the Boz application will be running.
- **Value:** `localhost`
- **Example Usage:**
  ```plaintext
  HOST=localhost
#### PORT
- **Description:** The port on which the Boz application will listen for incoming requests.
- **Value:** 8085
- **Example Usage:**
    ```plaintext
    PORT=8085
#### DB_USERNAME
- **Description:** The username used to connect to the PostgreSQL database.
- **Value:** boz
- **Example Usage:**
    ```plaintext
    DB_USERNAME=boz
#### DB_PASSWORD
- **Description:** The password used to authenticate the user connecting to the PostgreSQL database.
- **Value:** bozi
- **Example Usage:**
    ```plaintext
    DB_PASSWORD=bozi
#### DB_HOSTNAME
- **Description:** The hostname or IP address of the PostgreSQL database server.
- **Value:** localhost
- **Example Usage:**
    ```plaintext
    DB_HOSTNAME=localhost
#### DB_PORT
- **Description:** The port on which the PostgreSQL database server is listening.
- **Value:** 5432
- **Example Usage:**
    ```plaintext
    DB_PORT=5432
#### DB_DBNAME
- **Description:** The name of the PostgreSQL database to connect to.
- **Value:** boz
- **Example Usage:**
    ```plaintext
    DB_DBNAME=boz
#### MIGRATION_PATH
- **Description:** The path to the directory containing database migration files.
- **Value:** db/migration
- **Example Usage:**
    ```plaintext
    MIGRATION_PATH=db/migration
## Getting Started

To get started with the Boz application, follow these steps:

1. Set up your environment variables by creating a `.env` file. Example configurations can be found in `.env.example`.

2. Run the Boz application using the `main.go` file in the `cmd/boz` directory.

```bash
go run cmd/boz/main.go
```

Access the API routes as defined in the `routes/sheep.route.go` file.
## Notes
- The application uses Gin as the web framework.
The database is PostgreSQL, and queries are generated using SQLC.
- Configuration management is handled by Viper.
Feel free to explore and customize the code for your specific use case. 
- If you have any questions or encounter issues, please refer to the documentation of the respective packages used in this project.