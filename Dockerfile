ARG BASE_REGISTRY=zot.jamet.me/docker-hub/library

FROM ${BASE_REGISTRY}/golang:alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /bin/go-http-echo .

FROM ${BASE_REGISTRY}/alpine
COPY --from=build /bin/go-http-echo /bin/go-http-echo
ENTRYPOINT ["/bin/go-http-echo"]
EXPOSE 8080
