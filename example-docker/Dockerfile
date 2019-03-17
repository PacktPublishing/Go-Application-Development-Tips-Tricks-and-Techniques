FROM golang:1.11 as builder

COPY . /work
WORKDIR /work
RUN useradd app
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o hello

FROM scratch

COPY --from=builder /work/hello /usr/bin/hello
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/

USER app

ENTRYPOINT ["/usr/bin/hello"]]
