# Stage 1: Build the Go backend
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .


# # Expose ports for the frontend and backend
EXPOSE 8082

# Start Nginx and the Go application
CMD ["go","run", "main.go"]
