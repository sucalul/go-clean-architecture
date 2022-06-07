package main

import (
	"github.com/yuya0729/go-clean-architecture/cmd/app/driver"

	_ "github.com/lib/pq"
)

func main() {
	driver.Serve()
}
