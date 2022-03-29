FROM golang:1.18 as builder

RUN mkdir app
COPY . /app
WORKDIR /app/cmd

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o translatord

FROM gcr.io/distroless/static
COPY --from=builder /app/cmd/translatord /
CMD ["./translatord"]