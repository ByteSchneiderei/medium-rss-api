FROM golang:1.14.0 as builder
WORKDIR /go/src/github.com/ByteSchneiderei/medium-rss-api
COPY . .
RUN make linux 


FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/github.com/ByteSchneiderei/medium-rss-api/build/medium-rss-api /medium-rss-api

CMD ["/medium-rss-api"]
