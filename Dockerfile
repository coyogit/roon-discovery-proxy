FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go build -o roon-discovery-proxy main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/roon-discovery-proxy .
ENTRYPOINT ["./roon-discovery-proxy"]
