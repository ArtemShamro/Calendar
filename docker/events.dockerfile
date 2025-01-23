#build stage
FROM golang:1.23 AS builder
ARG SERVICE_NAME
WORKDIR $GOPATH/src/github.com/ArtemShamro/Calendar
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o app cmd/events/main.go

#final stage
FROM alpine:latest
WORKDIR /root/
RUN mkdir -p ./cmd/events
COPY --from=builder /go/src/github.com/ArtemShamro/Calendar .
# COPY --from=builder /go/src/github.com/irahardianto/monorepo-microservices/config/config.yaml ./config/
CMD ["./app"]

EXPOSE 8080

# docker build -t irahardianto/cinema-movies:v1 -f ./docker/booking.dockerfile --build-arg SERVICE_NAME=booking .