FROM golang:1.11 as builder

RUN useradd app

FROM scratch

COPY hello /usr/bin/hello
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/

USER app

ENTRYPOINT ["/usr/bin/hello"]]
