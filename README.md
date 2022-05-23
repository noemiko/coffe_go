# coffe_go
<!-- go mod init github.com/noemiko/coffe_go
go build main.go -->

Read about go
https://go.dev/doc/faq#Is_Go_an_object-oriented_language

This app solve first step of this kata:

https://simcap.github.io/coffeemachine/cm-first-iteration.html

How to build by docker

```bash
docker build -t apigopgm . --rm
docker run -p 4000:4000 --name apicontainer --rm apigopgm
```

how to build by docker-compose

```bash
docker-compose up --build
```

How to run tests
```bash
docker run --name apicontainer --rm apigopgm go test
```

How to test
```bash
curl -X POST -H "Content-Type: application/json" -d '{"command":"H:2:1"}' http://localhost:4000/drink

```


