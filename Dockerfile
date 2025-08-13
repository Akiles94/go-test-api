# Build stage
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata

# Create appuser
RUN adduser -D -g '' appuser

# Build argument for service path
ARG SERVICE_PATH

WORKDIR /build

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download
# Generate API documentation

# Copy entire project (needed for shared modules)
COPY . .

# Build the specific service binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo \
    -ldflags '-extldflags "-static"' \
    -o main ${SERVICE_PATH}/cmd/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates wget
WORKDIR /root/

# Copy the binary
COPY --from=builder /build/main .

# Default port (can be overridden)
EXPOSE 8080

# Run the binary
CMD ["./main"]