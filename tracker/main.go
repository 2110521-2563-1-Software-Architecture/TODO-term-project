package main

import (
	"tracker/routers"
	"tracker/services"
)

func main() {
	services.RegisterTrackerWithConsul()
	r := routers.SetupRouter()
	r.Run(services.Port())
}
