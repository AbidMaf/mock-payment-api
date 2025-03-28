# Payment Mock API

## Introduction

Simple payment API with go.

## Requirements

Before running the project, make sure to have these installed:

- [Go](https://go.dev/dl/) (>= 1.24)
- [Docker & Docker Compose](https://www.docker.com/) (optional, for deployment)

## Setup & Installation

### 1. Clone the Repository

```sh
git clone https://github.com/abidMaf/mock-payment-api.git
cd payment-mock-api
```

### 2. Setup Environment Variables

Create an `.env` file in the root directory and add the required environment variables:

```env
SERVER_PORT=your server port
JWT_SECRET=your secret jwt key
CUSTOMER_DB=location of customers storage file
MERCHANT_DB=location of merchants storage file
HISTORY_DB=location of history storage file
SESSION_DB=location of session storage file
```

or use `.env.example` as a template.

### 3. Install Dependencies

```sh
go mod tidy
```

## Running the Project

### Without Docker (Go Directly)

```sh
go run main.go
```

### With Docker

```sh
docker build -t bank-api .
docker run -p 8080:8080 bank-api
```

## Deployment

### Docker Compose

```sh
docker-compose up --build
```
