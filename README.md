# coffe_go
go mod init github.com/noemiko/coffe_go
go get -d "github.com/gin-gonic/gin"
go build main.go

docker build -t apigopgm . --rm
docker run -p 4000:4000 --name apicontainer --rm apigopgm

