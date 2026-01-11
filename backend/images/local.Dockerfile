FROM golang:1.25-alpine AS builder

WORKDIR /app
RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main

# ---------------------------------------------------------------- #

FROM golang:1.25-alpine

WORKDIR /app
# COPY --from=builder /app/main .
COPY --from=builder /go/bin/air /usr/local/bin/air
COPY --from=builder /app .
COPY .air.toml ./

EXPOSE 5000
CMD ["air", "-c", ".air.toml"]
