FROM cgr.dev/chainguard/go:latest as builder
ENV CGO_ENABLED=0
RUN mkdir -p /var/spool/peephole
WORKDIR /workspace
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build .

FROM ghcr.io/anchore/syft:latest AS sbomgen
COPY --from=builder /workspace/peephole /usr/sbin/peephole
RUN ["/syft", "--output", "spdx-json=/tmp/peephole.spdx.json", "/usr/sbin/peephole"]

FROM cgr.dev/chainguard/static:latest
EXPOSE 8080
WORKDIR /tmp
COPY --from=builder /workspace/peephole /usr/sbin/
COPY --from=builder /var/spool/peephole /var/spool/peephole
COPY --from=sbomgen /tmp/peephole.spdx.json /var/lib/db/sbom/peephole.spdx.json
CMD ["/usr/sbin/peephole"]
LABEL org.opencontainers.image.title="peephole"
LABEL org.opencontainers.image.description="SaltStack transactions tracker"
LABEL org.opencontainers.image.url="https://github.com/immobiliare/peephole"
LABEL org.opencontainers.image.source="https://github.com/immobiliare/peephole"
LABEL org.opencontainers.image.licenses="MIT"
LABEL io.containers.autoupdate=registry
