# Use Golang image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy Go source
COPY . .

# Build the Go app
RUN go mod init golangmagicball && go build -o main .

# Expose port
EXPOSE 8080

# Run it
CMD ["./main"]
