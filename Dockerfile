FROM golang:1.20-alpine as builder

# Create a non-root user
RUN adduser --disabled-password --gecos "" ports-user

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd cmd
COPY internal internal

RUN chown -R ports-user /app
USER ports-user

RUN go build -o ./build/ports-service ./cmd/ports
RUN go build -o ./build/webapp ./cmd/webapp


FROM alpine:3.17 as ports

WORKDIR /app
COPY --from=builder /app/build/ports-service .

CMD ["./ports-service"]

FROM alpine:3.17 as webapp

WORKDIR /app
COPY --from=builder /app/build/webapp .

CMD ["./webapp"]

