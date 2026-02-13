# simple-go-app

A simple Go application that serves a static site using [Fiber](https://gofiber.io/).

## Description

This project demonstrates a basic Go web server setup using the Fiber framework. It serves static HTML files and resources, making it a good starting point for simple web applications or static sites.

## Prerequisites

- [Go](https://go.dev/dl/) 1.25 or higher
- [Docker](https://www.docker.com/products/docker-desktop/) (optional, for containerization)

## Usage

### Run Locally

To run the application locally:

```bash
go run main.go
```

The server will start on port `3000` by default. You can access it at `http://localhost:3000`.

To specify a different port:

```bash
PORT=8080 go run main.go
```

### Run Tests

To run the unit tests:

```bash
go test -v ./...
```

### Build

To build the application binary:

```bash
go build -o simple-go-app main.go
```

## Docker

### Build Image

To build the Docker image:

```bash
docker build -t simple-go-app .
```

### Run Container

To run the application in a Docker container:

```bash
docker run -p 8080:8080 -e PORT=8080 simple-go-app
```

The application will be accessible at `http://localhost:8080`.

## Deployment

### Google App Engine

To deploy to Google App Engine:

1. Install and setup the [Google Cloud SDK](https://cloud.google.com/sdk/docs/install).
2. Run the deployment command:

```bash
gcloud app deploy
```

## CI/CD

This project uses GitHub Actions for Continuous Integration. The workflow is defined in `.github/workflows/go.yml` and triggers on push and pull requests to the `master` branch. It performs the following steps:

1. Checks out the code.
2. Sets up the Go environment.
3. Builds the project to ensure compilation success.
