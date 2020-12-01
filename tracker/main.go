package main

import (
	"tracker/routers"
	"tracker/services"
	"tracker/services/consul"
)

func main() {
	consul.RegisterTrackerWithConsul()
	r := routers.SetupRouter()
	r.Run(services.Port())
}
