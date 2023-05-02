# Use the official Golang image as the base image
FROM golang:1.17 as build

# Set the working directory
WORKDIR /app

# Enable Go modules explicitly
ENV GO111MODULE=on

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the application
RUN go build -o chorechamps .

# Use a minimal image for deployment
FROM gcr.io/distroless/base

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/chorechamps /app/

# Expose the application port
EXPOSE 8282

# Set the entrypoint for the container
CMD ["./chorechamps"]