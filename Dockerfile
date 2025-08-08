FROM golang:1.24-alpine

# Set working directory inside the container
WORKDIR /usr/app

# Pre-fetch dependencies (speeds up dev iterations)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application (useful if running image without bind mount)
COPY . .

# Expose API port
EXPOSE 8080

# Use Air for hot-reload in dev; docker-compose will also override this CMD
CMD ["go", "run", "github.com/air-verse/air"]

