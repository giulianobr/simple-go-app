# Stage 1: Build the Go application
FROM golang:1.25-alpine AS builder

# Set necessary environment variables
ENV CGO_ENABLED=0
ENV GOOS=linux

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies.
# This leverages Docker's layer caching.
COPY go.mod ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app, creating a static binary
RUN go build -o /simple-go-app -ldflags="-w -s" .

# Stage 2: Create the final, minimal, non-root image
FROM gcr.io/distroless/base AS final

WORKDIR /

# Copy the static binary from the builder stage
COPY --from=builder /simple-go-app /simple-go-app

# Copy the static web assets
COPY --from=builder /app/web /web

# Expose port 8080 (the default for Gin)
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/simple-go-app"]
