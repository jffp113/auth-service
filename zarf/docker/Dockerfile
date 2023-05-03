FROM golang:1.20.4 AS builder

WORKDIR /builder
COPY . .

RUN go build --ldflags '-linkmode external -extldflags "-static"' ./app/service/auth-api/main.go


FROM alpine:3.17


WORKDIR service
COPY --from=builder /builder/main ./auth-api

RUN chmod +x ./auth-api

ENTRYPOINT ["./auth-api"]