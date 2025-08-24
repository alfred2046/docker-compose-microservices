# Stage 1: build go
FROM golang:1.25 AS builder

WORKDIR /app

COPY . .
RUN go mod tidy
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: create runtime
FROM alpine:latest

WORKDIR /opt/msproductapi/
COPY --from=builder /app/main .

EXPOSE 8082
CMD ["./main"]