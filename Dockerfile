FROM golang:1.23.0-alpine

WORKDIR /app

COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o goshort .

# Expose the port the application runs on
EXPOSE 8080

# Run the application
CMD ["./goshort"]
