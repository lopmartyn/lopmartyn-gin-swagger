## Prerequisite

- Bun, ORM framework
`$ go get github.com/uptrace/bun`
- Gin Web Framework
  `$ go get -u github.com/gin-gonic/gin`
- PostgreSQL
  `$ go get github.com/uptrace/bun/dialect/pgdialect`
`$ go get github.com/uptrace/bun/driver/pgdriver`
- Swagger `$ go get -u github.com/go-swagger/go-swagger/cmd/swagger`
- logrus `$ go get -u github.com/sirupsen/logrus` 
- Docker Desktop
- Goland IDE(if you have licence)

## Start project
- Go to resource/docker
- Start docker `docker compose up`
- Go to main.go (default port is 9000)
- cmd `go run main` 
- you can access to swagger document at http://localhost:9000/swagger/index.html