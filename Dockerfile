# Use a lightweight Go base image
FROM golang:1.22

WORKDIR /app

COPY ./life-server/ .

WORKDIR /app/cmd

RUN go build main.go

EXPOSE 8080

# Set the entry point for the container
CMD ["./main"]
