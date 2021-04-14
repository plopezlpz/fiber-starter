# Template go-fiber project structure

Opinionated template for a Rest API microservice using [gofiber](https://github.com/gofiber/fiber) web framework.

The architecture is intended to look like this:

![architecture](https://raw.githubusercontent.com/plopezlpz/fiber-starter-doc/master/assets/images/architecture-go-fiber-bg.png)

Folder structure:
```
.
├── Makefile
├── README.md
├── config
│   └── vars
│       └── vars.go
├── domain
│   └── domain.go
├── dto
│   └── doc.go
├── go.mod
├── go.sum
├── handler
│   └── health.go
├── main.go
├── model
│   └── doc.go
├── rest
│   ├── rest.go
│   └── router.go
└── service
    └── service.go
```


## Quick Start

Copy the sample `.env-sample` file to `.env`
```bash
cp .env-sample .env
```
Run
```bash
make run
```
Test
```bash
curl localhost:8888/health
```
