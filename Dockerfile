ARG BASE_REGISTRY=zot.jamet.me/docker-hub/library
ARG GO_VERSION=1.27.3
ARG ALPINE_VERSION=3.24

FROM ${BASE_REGISTRY}/golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /bin/go-http-echo .

FROM ${BASE_REGISTRY}/alpine:${ALPINE_VERSION}
COPY --from=build /bin/go-http-echo /bin/go-http-echo
ENTRYPOINT ["/bin/go-http-echo"]
EXPOSE 8080
