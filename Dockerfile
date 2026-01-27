# ===== BUILD STAGE =====
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

# copy dependency files first (for cache)
COPY go.mod go.sum ./
RUN go mod download

# copy source
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o coupon-system ./cmd/api

# ===== RUNTIME STAGE =====
FROM alpine:3.21

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/coupon-system .

EXPOSE 8080

CMD ["./coupon-system"]
