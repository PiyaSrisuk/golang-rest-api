FROM golang:1.24.2-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o app .

FROM alpine:latest

COPY --from=builder /app/app /app/

CMD ["/app/app"]
