FROM golang:1.17 as builder
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
WORKDIR /app
COPY --from=builder /app/peephole .
COPY --from=builder /app/kiosk/assets kiosk/assets
CMD ["./peephole", "-c", "./configuration.yml"]