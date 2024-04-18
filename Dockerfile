# Use an official Golang runtime as a parent image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace.
COPY go.mod ./
COPY *.go ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
RUN go build -o myapp

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]
