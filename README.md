# coffe_go
go mod init github.com/noemiko/coffe_go
go build main.go

How to build
```
docker build -t apigopgm . --rm
docker run -p 4000:4000 --name apicontainer --rm apigopgm
```

How to run tests
```
docker run --name apicontainer --rm apigopgm go test
```

How to test
```
curl -X POST -H "Content-Type: application/json" -d '{"command":"H:2:1"}' http://localhost:4000/drink

```

