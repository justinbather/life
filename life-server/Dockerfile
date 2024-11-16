# Use a lightweight Go base image
FROM golang:1.22

# Set the working directory in the container
WORKDIR /app

# Copy only the service-specific files
COPY . .

# Build the Go application (assuming `main.go` is in `cmd/`)
WORKDIR /app/cmd
RUN go build main.go

# Set the entry point for the container
CMD ["./main"]
