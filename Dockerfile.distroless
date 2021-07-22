FROM golang:1.16.6-alpine as build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

FROM gcr.io/distroless/static:nonroot

COPY --from=build /app/geolocation /usr/bin/geolocation

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=2 CMD curl -f http://localhost/ || exit 1

ENTRYPOINT ["/usr/bin/geolocation"]
