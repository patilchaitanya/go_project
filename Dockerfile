# Use the official Go image as the base image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the Go server source code into the container
COPY . .

# Initialize modules and download dependencies
RUN go mod init hello.go && go mod tidy

# Build the Go server binary
RUN go build -o main .

# Expose the port the server will run on
EXPOSE 8000

# Command to run the Go server
CMD ["./main"]
