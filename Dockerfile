# Build Stage
FROM golang:1.22.0-alpine as builder

# Add maintainer
LABEL maintainer="Srijil V<srijilvijayan@gmail.com>"


# Installing git for fetching dependencies.
RUN apk update && apk add --no-cache git

# RUN git config --global http.sslVerify "false"

# Set the working directory inside the container
WORKDIR /app

# Copy go modules and install dependencies
COPY go.mod go.sum ./

RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/cli/openapi/main.go

# Final Runtime Stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/app .

# Expose the application port (change as needed)
EXPOSE 8080

# Command to run the executable
CMD ["./app"]