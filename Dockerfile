# Stage 1: build
FROM golang:1.22 AS builder

WORKDIR /app

# Set GOPROXY supaya dependency bisa didownload
ENV GOPROXY=https://proxy.golang.org,direct

# Copy go.mod & go.sum dulu biar cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code
COPY . .

# Build binary
RUN go build -o server ./main.go

# Stage 2: runtime
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# Copy binary & assets dari builder
COPY --from=builder /app/server .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

# Expose port
EXPOSE 8080

CMD ["./server"]
