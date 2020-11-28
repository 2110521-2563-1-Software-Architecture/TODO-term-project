package main

import (
	"peer/routers"
	"peer/services/consul"
)

func main() {
	consul.RegisterServiceWithConsul()
	r := routers.SetupRouter()
	r.Run(":9000")
}
