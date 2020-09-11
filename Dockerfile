FROM alpine:latest as alpine

RUN apk add -U --no-cache ca-certificates

FROM scratch
WORKDIR /
COPY /main /main
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/main"]