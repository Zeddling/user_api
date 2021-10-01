package main

import (
	"flag"
	"fmt"

	"github.com/Zeddling/user/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 8080, "port number")
	engine := routes.StartGin(port)
	engine.Run(fmt.Sprintf(":%d", port))
}
