FROM alpine as builder
RUN apk add --allow-untrusted --update --no-cache curl ca-certificates
WORKDIR /
RUN curl -fsSL github.com/metrico/iox-builder/releases/latest/download/influxdb3 -O && chmod +x influxdb3

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /influxdb3 /influxdb
EXPOSE 8181
CMD ["/influxdb", "serve", "--writer-id=iox", "--object-store", "file", "--data-dir", "/data"]
