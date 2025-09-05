# syntax=docker/dockerfile:1
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o payments-api main.go

FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/payments-api /payments-api
EXPOSE 8081
ENTRYPOINT ["/payments-api"]
