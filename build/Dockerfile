FROM golang:1.18 as build

RUN mkdir app
COPY . /app
WORKDIR /app/cmd

RUN go build -o translatord

FROM gcr.io/distroless/base-debian11
COPY --from=build /app/cmd/translatord /
CMD ["./translatord"]