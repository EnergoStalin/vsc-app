package providers

import "github.com/EnergoStalin/vsc-app/pkg/vsc-providers/responces"

type SearchProvider interface {
	GetName() string
	SearchByInn(inn string) responces.SearchResponce
	SearchByName(name string) responces.SearchResponce
}

var Providers = []SearchProvider{
	&ZaChestnyiBiznesProvider{},
}
