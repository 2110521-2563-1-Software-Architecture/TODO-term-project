package services

import (
	"fmt"
	"log"
	"os"
	"strconv"

	consulapi "github.com/hashicorp/consul/api"
)

func RegisterTrackerWithConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}

	registration := new(consulapi.AgentServiceRegistration)

	id := os.Getenv("REGISTRATION_ID")
	registration.ID = "tracker-service-" + id
	registration.Name = "tracker-service"
	address := Hostname()
	registration.Address = address
	p, err := strconv.Atoi(Port()[1:len(Port())])
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

func LookupPeersWithConsul() ([]string, error) {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		return []string{}, err
	}

	services, err := consul.Agent().ServicesWithFilter("Service == \"peer-service\"")
	if err != nil {
		return []string{}, err
	}

	var result []string

	for _, val := range services {
		addr := fmt.Sprintf("http://%s:%v", val.Address, val.Port)
		result = append(result, addr)
	}
	return result, nil
}
