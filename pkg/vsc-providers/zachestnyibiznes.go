package providers

import "github.com/EnergoStalin/vsc-app/pkg/vsc-providers/responces"

type ZaChestnyiBiznesProvider struct {
}

func (z *ZaChestnyiBiznesProvider) SearchByInn(inn string) responces.SearchResponce {
	return responces.SearchResponce{}
}

func (z *ZaChestnyiBiznesProvider) SearchByName(inn string) responces.SearchResponce {
	return responces.SearchResponce{}
}

func (z *ZaChestnyiBiznesProvider) GetName() string {
	return "ZaChestnyiBiznes"
}
