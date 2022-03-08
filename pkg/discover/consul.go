package discover

import (
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/wagaru/ticket/pkg/config"
)

func NewConsulClient() {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = config.ConsulHost
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		config.Logger.Log("err", err)
	}
	DiscoverClient = consul.NewClient(consulClient)
}
