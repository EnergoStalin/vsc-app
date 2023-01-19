package providers

import (
	"context"

	"github.com/EnergoStalin/vsc-app/pkg/vsc-providers/responces"
)

type SearchProvider interface {
	GetName() string
	GetUrl() string
	Search(name string, ctx context.Context) (responces.SearchResponce, error)
}

var Providers = []SearchProvider{
	NewZaChestnyiBiznesProvider(),
}
