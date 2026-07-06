FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Force rebuild by setting a build timestamp
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X main.BuildTime=$(date)" -o much-to-do ./cmd/api/main.go

FROM alpine:3.20

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/much-to-do .

EXPOSE 8080

ENTRYPOINT ["./much-to-do"]
