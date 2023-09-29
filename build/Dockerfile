FROM golang:1.21.1-bullseye as builder

WORKDIR /build

# Load dependencies
COPY go.mod go.sum ./
RUN go mod download

# COPY the source code as the last step
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -a -o /placeapi cmd/place-api/main.go

FROM debian:stable

WORKDIR /app

COPY --from=builder /placeapi .
COPY --from=builder /build/config/config.yml .

ENTRYPOINT [ "/app/placeapi" ]
