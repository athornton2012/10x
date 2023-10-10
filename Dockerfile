FROM golang:1.20 as build

ENV CGO_ENABLED=0

WORKDIR /app

# Put dependencies in their own layer so they are cached.
COPY go.mod go[.]sum ./
RUN go mod download

# Then install the rest of the code
COPY . .

RUN go build -o . ./...

FROM build AS test
RUN go test -v ./...


FROM gcr.io/distroless/base-debian11 AS run

WORKDIR /app
COPY --from=build /app/10x .
COPY --from=build /app/fixtures/seattle-weather.csv .

EXPOSE 8080

ENTRYPOINT ["/app/10x", "seattle-weather.csv"]