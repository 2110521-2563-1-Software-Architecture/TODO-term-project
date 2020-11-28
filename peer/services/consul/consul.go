package consul

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"peer/services"

	consulapi "github.com/hashicorp/consul/api"
)

func RegisterServiceWithConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}

	registration := new(consulapi.AgentServiceRegistration)

	id := os.Getenv("REGISTRATION_ID")
	registration.ID = "peer-service-" + id
	registration.Name = "peer-service"
	address := services.Hostname()
	registration.Address = address
	p, err := strconv.Atoi(services.Port()[1:len(services.Port())])
	if err != nil {
		log.Fatalln(err)
	}
	registration.Port = p
	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/healthcheck", address, p)
	registration.Check.Interval = "5s"
	registration.Check.Timeout = "3s"
	consul.Agent().ServiceRegister(registration)
}
