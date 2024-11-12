# Used an official Go runtime as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copied the Go file into the container at /app
COPY primenumbers.go .

# Built the Go app
RUN go build -o primenumbers primenumbers.go

# EXPOSE 8080

# Run the binary
CMD ["./primenumbers"]
