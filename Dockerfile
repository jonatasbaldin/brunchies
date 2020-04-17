FROM golang:1.13-alpine as builder

WORKDIR /go/src/app
COPY brunchies/* /go/src/app/

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o /go/bin/brunchies

FROM scratch

COPY brunchies/index.html /
COPY --from=builder /go/bin/brunchies /brunchies
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["/brunchies"]