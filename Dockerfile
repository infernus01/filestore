FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY pkg/ pkg/
COPY cmd/ cmd/

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -o server cmd/file-server/main.go

# Runtime Stage
FROM scratch

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server /

ENTRYPOINT ["/server"]