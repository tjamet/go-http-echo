# go-http-echo
A simple service that logs all the requests and responds with the original request content

# run

```bash
go get -u github.com/tjamet/go-http-echo
go-http-echo
curl localhost:8080/hello -d world
```

```bash
docker run -P --rm --name http-echo tjamet/http-echo
curl $(docker port http-echo 8080)/hello -d world
```
