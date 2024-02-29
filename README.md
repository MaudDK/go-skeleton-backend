# GoLang Backend Skeleton Project

This repository serves as a skeleton for kickstarting a backend project using GoLang. It comes pre-configured with:

- PostgreSQL database connectivity
- Example routes for creating an account
- RESTful API structure
- SQL migrations support

## Docker

To run this project using Docker, follow these steps:

1. Ensure you have Docker installed on your machine.

2. Create a `.env` file in the project root directory with the following content:

    ```
    # PostgreSQL environment variables
    POSTGRES_DB=postgres
    POSTGRES_USER=postgres
    POSTGRES_PASSWORD=postgres
    
    # Backend Environment Variables
    STAGE=dev
    PORT=4000
    ```

3. Build and run the Docker containers using Docker Compose:

    ```bash
    docker-compose up --build
    ```

4. The server will start inside a Docker container and will be accessible at `http://localhost:port`.

## Configuration (Without Docker)
1. Ensure that PostgreSQL is running.
2. Configure the PostgreSQL connection enviorment variable parameters.
    ```bash
    export POSTGRES_DB=database_name
    export POSTGRES_USER=username
    export POSTGRES_PASSWORD=password
    export STAGE=Development
    export PORT=PORT
    ```

## Prerequisites

- GoLang installed on your machine.
- PostgreSQL installed and running locally or accessible via a network.
- `go get` for dependency management.

## Installation

1. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/MaudDK/go-skeleton-backend.git
    ```

2. Navigate into the project directory:

    ```bash
    cd go-skeleton-backend
    cd backend
    ```

3. Install dependencies:

    ```bash
    go get -d -v ./...
    go mod tidy
    ```
## Usage

1. Run the migrations to set up the database schema:

    ```bash
    go build -o ./bin/app ./cmd/main.go
    ```

2. Start the server:

    ```bash
    "./bin/app"
    ```

3. The server will start at `http://localhost:PORT`. You can use tools like Postman or Insomnia to interact with the API.

## API Endpoints

- **GET /api/v1/health**: Check the health status of the application.
- **POST /api/v1/accounts**: Create a new account.
- **GET /api/v1/accounts/:id**: Retrieve account details by ID.