# Use the official Golang image from the Docker Hub
FROM golang:1.23.4-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download dependencies and build the application
RUN go mod tidy
RUN go build -o main .

# Command to run the executable
CMD ["./main"]