# Template go-fiber project structure

Opinionated template for a Rest API microservice using [gofiber](https://github.com/gofiber/fiber) web framework.

The architecture is intended to look like this:

![architecture](https://raw.githubusercontent.com/plopezlpz/fiber-starter-doc/master/assets/images/architecture-go-fiber-bg.png)


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
