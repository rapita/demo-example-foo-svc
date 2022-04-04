FROM golang:1.17-bullseye as build

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

ADD . .

RUN CGO_ENABLED=0 go build -o ./app ./cmd/main.go

# "Distroless" images contain only your application and its runtime dependencies.
# They do not contain package managers, shells or any other programs you would expect to find in a standard Linux distribution.
#
# See: https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian10

COPY --from=build /go/src/app/app /bin/app

USER nonroot:nonroot

ENTRYPOINT ["/bin/app"]
