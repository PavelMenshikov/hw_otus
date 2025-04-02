FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o financial_planner

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/financial_planner .
CMD ["./financial_planner"]