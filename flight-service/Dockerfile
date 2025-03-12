FROM golang:1.23-bookworm AS builder
# FROM golang:1.21-bookworm AS builder

WORKDIR /app

# Copy go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the entire source code
COPY . .
COPY /credentials/credentials.json /app/credentials/credentials.json
COPY conf /app/conf
# Set environment variables (without credentials.json)
ENV PROJECT_ID=microservices-451611 \
    INSTANCE_ID=microservice \
    DATABASE_ID=microservice \
    GOOGLE_APPLICATION_CREDENTIALS="/app/credentials/credentials.json"

# Build the Go application
RUN chmod 644 /app/credentials/credentials.json
# RUN go build -o main .                                                                                        
RUN CGO_ENABLED=0 go build -o main .

# Final Image (distroless for minimal size)
FROM gcr.io/distroless/base-debian12

WORKDIR /app




# Copy compiled binary and .env file from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .  
# Copy the necessary config files
COPY --from=builder /app/credentials/credentials.json /app/credentials/credentials.json
COPY --from=builder /app/conf /app/conf

# Set environment variables
ENV GOOGLE_APPLICATION_CREDENTIALS="/app/credentials/credentials.json"
ENV CHASSIS_CONF_DIR="/app/conf"


# Expose the service port
EXPOSE 8080

# Run the application
# CMD ["./main"]

# RUN chmod +x /app/main

# Run the application
CMD ["/app/main"]