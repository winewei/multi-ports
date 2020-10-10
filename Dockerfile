FROM golang:1.14-alpine3.12 AS builder

WORKDIR /srv
COPY . .

RUN go build -o multi-ports multi-ports.go

FROM alpine

WORKDIR /srv
ENV PORTS="80,8080"

COPY --from=builder /srv/multi-ports multi-ports

EXPOSE 8080

CMD ["/srv/multi-ports"]
