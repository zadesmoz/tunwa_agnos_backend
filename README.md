# Tunwa Agnos Backend

## Overview

This repository contains the backend implementation of the Tunwa Agnos project, which includes a password strength recommendation API. The backend is built using Go with the Gin framework, Nginx, PostgreSQL, and Docker Compose.

## Features

- Password strength validation and recommendation API.
- Logging of requests and responses.
- Unit tests for core functionalities.

## Prerequisites

- [Go](https://golang.org/doc/install) 1.22.5 or higher
- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)
- [PostgreSQL](https://www.postgresql.org/download/)

## Installation

### 1. Build and run the application with Docker Compose

```bash
docker compose -f "docker-compose.yml" up -d --build
```
### 2. Running Unit Tests
```bash
go test tunwa/utils
```
### Run tests with VSCode
If you're using Visual Studio Code, you can also run the tests using the provided launch configuration.

- Open the command palette (Ctrl + Shift + P).
- Select Tasks: Run Test Task.
- Choose the `Launch Unit Tests Function`.

### 3. Access the API
Once the application is running, you can access the API at http://localhost:3000.
You can test API using tools like curl or Postman:

- POST `/api/strong_password_steps`
  - Request Body:
    ```json
    {"init_password": "aA1"}
    ```
