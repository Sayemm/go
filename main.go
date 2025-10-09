package main

import (
	"ecommerce/cmd"
)

func main() {
	cmd.Serve()
}

/*
Infrastructure
==============
- DB
- redis
- rabbitmq
- kafka
- file storage

db
===
- application connects with db and app gets a client
- client will store/delete... data from db
- database connection - library sqlx/sqlc/ORM (gorm)
*/
