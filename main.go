package main

import (
	"fmt"

	"github.com/hypebeast/go-osc/osc"
)

func main() {
	fmt.Println("Hello sicc")
}

func newClient(addr string, port int) client {
	return client{
		oscClient: osc.NewClient(addr, port),
	}
}
