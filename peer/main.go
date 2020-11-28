package main

import (
	"peer/routers"
	"peer/services"
	"peer/services/consul"
)

func main() {
	consul.RegisterServiceWithConsul()
	r := routers.SetupRouter()
	r.Run(services.Port())
}
