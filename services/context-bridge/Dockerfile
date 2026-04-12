# ---- Build stage ----
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/context-bridge ./cmd/main/...

# ---- Run stage ----
FROM alpine:3.19
RUN addgroup -S app && adduser -S app -G app
USER app
COPY --from=builder /bin/context-bridge /bin/context-bridge
EXPOSE 8081
HEALTHCHECK --interval=10s --timeout=3s CMD wget -qO- http://localhost:8081/healthz || exit 1
ENTRYPOINT ["/bin/context-bridge"]
