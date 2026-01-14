package main

import (
	"ecommerce/cmd"
	"sync"
)

var cnt int64
var mu sync.Mutex

func main() {
	cmd.Serve()
}
