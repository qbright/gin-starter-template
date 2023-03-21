# Go Gin Starter Template
> A starter project with Golang, Gin and Gorm and some gin middleware like session , cors, logger

> support development hot reload use air



### structure

```
|-- .air.toml
|-- .gitignore
|-- Makefile
|-- README.md
|-- go.mod
|-- go.sum
|-- main.go
|-- config
|   |-- config.go
|   |-- development.yaml
|   |-- production.yaml
|   |-- test.yaml
|-- controllers
|   |-- health.go
|   |-- user.go
|-- db
|   |-- init.go
|   |-- dbutil
|       |-- migrate.go
|-- middlewares
|   |-- auth.go
|-- modules
|   |-- base.go
|   |-- user.go
|-- server
|   |-- router.go
|   |-- server.go
|-- utils
    |-- logger.go

```

## Installation

```sh
make deps
```

## Usage example

`curl http://localhost:8080/health`

## Development setup

Running mysql on Docker Image: mysql:5.7

start development

`make run`



