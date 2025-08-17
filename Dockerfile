FROM golang:1.24.5-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o addon .

FROM alpine:latest
COPY --from=builder /app/addon /addon
EXPOSE 8080
ENTRYPOINT ["/addon"]
