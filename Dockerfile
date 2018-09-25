FROM golang:alpine AS build
COPY *.go /go/src/github.com/tjamet/go-http-echo/
RUN go build -o /bin/go-http-echo github.com/tjamet/go-http-echo

FROM alpine
COPY --from=build /bin/go-http-echo /bin/go-http-echo
ENTRYPOINT ["/bin/go-http-echo"]
EXPOSE 8080
