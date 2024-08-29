FROM golang:1.23-bookworm as builder

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build

FROM debian:bookworm-slim

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.

COPY --from=builder /app/hello /app/hello

# Run the web service on container startup.

CMD ["/app/hello"]



