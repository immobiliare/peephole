FROM golang:1.17 as builder
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /workspace
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go install github.com/gobuffalo/packr/packr@latest \
 && make build

FROM debian:buster
EXPOSE 8080
WORKDIR /tmp
COPY --from=builder /workspace/peephole /usr/sbin/
RUN mkdir -p /var/spool/peephole
CMD ["/usr/sbin/peephole"]
LABEL org.opencontainers.image.source=https://github.com/immobiliare/peephole
