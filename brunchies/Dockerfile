FROM golang:1.13 as build

RUN mkdir -p /go/src/handler
WORKDIR /go/src/handler
COPY . .

RUN CGO_ENABLED=0 GOOS=linux \
    go build --ldflags "-s -w" -a -installsuffix cgo -o handler . && \
    go test $(go list ./... | grep -v /vendor/) -cover

FROM alpine:3.9
# Add non root user and certs
RUN apk --no-cache add ca-certificates \
    && addgroup -S app && adduser -S -g app app \
    && mkdir -p /home/app \
    && chown app /home/app

WORKDIR /home/app

COPY --from=build /go/src/handler/handler .
COPY --from=build /go/src/handler/index.html .

RUN chown -R app /home/app

RUN touch /tmp/.lock #  Write a health check for OpenFaaS here or in the HTTP server start-up

USER app

CMD ["/home/app/handler"]