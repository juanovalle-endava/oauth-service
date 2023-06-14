# Build stage
FROM golang:alpine AS builder
# Set necessary environment variables needed for the image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GIN_MODE=release

# Move to working directory /build
WORKDIR /build

# Copy and download dependencies using go mod
COPY go.mod go.sum ./

RUN go mod download

# Copy the code into the container
COPY . .

# Go tests before building the app
# RUN go test ./...

WORKDIR /build/cmd

# Build the application
RUN go build -o main .

# Use a minimal Alpine image as the base image for the final image
FROM alpine:latest

# Move to /app directory as the place for the resulting binary folder
WORKDIR /app

# Copy binary from build to main folder
COPY --from=builder /build/cmd/main .
COPY --from=builder /build/config.yaml .
COPY --from=builder /build/cert ./cert/



# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["/app/main"]
