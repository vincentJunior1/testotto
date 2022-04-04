# Golang Gin GORM starter

This repository contains a starter template for rapid microservice development
with Go. It uses [Gin](https://github.com/gin-gonic/gin) and
[GORM](https://gorm.io).

## Installation
* Get the repository from GitHub
``` bash
git clone https://github.com/vincentJunior1/test-kriya.git
```
* Install dependencies
``` bash
go get
```

* Run app
```
go run main.go
```

## Things to consider
* Rename all instances of `https://bitbucket.org/vincent265/testmajoo/` to your package
* Switch from the default database driver (mysql) to sqlite, postgresql, ...

## Env variables

* PORT (Default: `8080`)
* GIN_MODE (Default: `debug`)
* DATABASE_HOST (Default: `localhost:3306`)
* DATABASE_NAME (Default: `testing`)
