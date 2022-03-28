FROM golang:1.18 as build

RUN mkdir app
COPY . /app
WORKDIR /app/cmd

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o translatord

FROM gcr.io/distroless/base-debian11
COPY --from=build /app/cmd/translatord /
CMD ["./translatord"]