## go-fiber-todo

A simple TODO application with Golang, Fiber, MongoDB and HTML Template engine

Stack: [Fiber](https://gofiber.io/), [Golang](https://golang.org/), [MongoDB](https://www.mongodb.com/), [HTML/Template](https://golang.org/pkg/html/template/)

DEV: http://localhost:5511

STAGE: https://golang-todo.herokuapp.com

### Deploy

Golang **1.15.X** is required

```shell script
git clone https://github.com/peterdee/go-fiber-todo
cd ./go-fiber-todo
```

### Environment variables

The `.env` file is required, see the [.env.example](.env.example) for details

### Launch (local)

```shell script
go run ./
```

Can be used with [AIR](https://github.com/cosmtrek/air), see the [run.sh](run.sh) file for details

### Heroku

The `stage` branch of the project is deployed to Heroku automatically
