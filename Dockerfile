FROM golang:alpine AS builder

WORKDIR /var/app

COPY . .

RUN go build cmd/weatherapi/main.go

FROM scratch

ARG API_KEY

WORKDIR /var/app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /var/app/.env .
COPY --from=builder /var/app/main .

EXPOSE 8080

ENTRYPOINT [ "./main" ]