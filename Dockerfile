# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o docker-ss-node ./cmd/simple-service/main.go

EXPOSE 2327

# Run
CMD ["./docker-ss-node"]
