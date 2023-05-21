FROM alpine as builder
RUN apk add --allow-untrusted --update --no-cache curl ca-certificates
WORKDIR /
RUN curl -fsSL github.com/metrico/iox-builder/releases/latest/download/influxdb_iox -O && chmod +x influxdb_iox

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /influxdb_iox /influxdb_iox
EXPOSE 8080 8082
CMD ["/influxdb_iox", "run"]
