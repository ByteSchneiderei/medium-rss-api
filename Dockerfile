FROM golang:1.14.0 as builder
WORKDIR /go/src/byteschneiderei.com/intern/medium/
COPY . .
RUN make linux 


FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/byteschneiderei.com/intern/medium/build/medium-rss-api /medium-rss-api

CMD ["/medium-rss-api"]
