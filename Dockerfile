FROM golang:1.15 as builder
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make build

FROM debian:buster
EXPOSE 8080
WORKDIR /app/
COPY --from=builder /app/peephole .
CMD ["./peephole -c /app/configuration.yml"]
