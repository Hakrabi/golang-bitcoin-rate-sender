FROM golang:1.17-alpine as build-stage

RUN apk --no-cache add ca-certificates

WORKDIR /go/golang-bitcoin-rate-sender

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /golang-bitcoin-rate-sender .

FROM scratch

COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build-stage /golang-bitcoin-rate-sender /golang-bitcoin-rate-sender

EXPOSE 8080

ENTRYPOINT ["/golang-bitcoin-rate-sender"]