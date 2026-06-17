# Stage 1: Build
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
# CGO_ENABLED=0 ensures static linking. -ldflags="-w -s" strips debug info to shrink binary size.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/header-inspection ./cmd

# Stage 2: Final minimal image
# Distroless is smaller than Alpine and contains no shell, reducing attack surface
FROM gcr.io/distroless/static-debian12:nonroot

# Copy the static binary
COPY --from=builder /app/header-inspection /app/header-inspection

EXPOSE 8080
# Run as the built-in nonroot user for better security
USER nonroot:nonroot

CMD ["/app/header-inspection"]