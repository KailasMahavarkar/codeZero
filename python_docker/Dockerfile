# Start with the Alpine Linux base image
FROM alpine:latest

# Install g++, Node.js, Go, and other necessary packages
RUN apk update && \
    apk add --no-cache g++ make nodejs npm go

# Set the working directory
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Expose the port the app runs on
EXPOSE 8080

# Start the Go server
CMD ["go", "run", "main.go"]
