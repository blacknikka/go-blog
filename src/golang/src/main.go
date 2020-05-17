package main

import (
	"github.com/blacknikka/go-blog/infra"
	"github.com/blacknikka/go-blog/infra/db/postgresql"
)

func main() {
	defer postgresql.CloseConn()

	infra.Router.Run()
}
