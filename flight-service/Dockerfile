# Build Stage
FROM golang:1.23-bookworm AS builder

WORKDIR /app

# Copy go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the entire source code
COPY . .

# Copy Google Cloud credentials
COPY /credentials/credentials.json /app/credentials.json

# Set environment variables
ENV PROJECT_ID=microservices-451611 \
    INSTANCE_ID=microservice \
    DATABASE_ID=microservice \
    GOOGLE_APPLICATION_CREDENTIALS="/app/credentials.json"

# Ensure credentials are readable
RUN chmod 644 /app/credentials.json

# Build the Go application
RUN go build -o main .

# Final Image (distroless for minimal size)
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# Copy compiled binary and credentials from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/credentials.json .

# Set the environment variable for Google Cloud authentication
ENV GOOGLE_APPLICATION_CREDENTIALS="/app/credentials.json"

EXPOSE 8080

# Run the application
CMD ["/app/main"]
