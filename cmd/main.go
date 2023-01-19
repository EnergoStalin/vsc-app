package main

import (
	"fmt"

	providers "github.com/EnergoStalin/vsc-app/pkg/vsc-providers"
)

func main() {
	for _, provider := range providers.Providers {
		for i := 0; i < 300; i++ {
			res, err := provider.Search("1-ая Внедренческая Компания", nil)
			fmt.Printf("%d %s %+v %s\n", i, provider.GetName(), res, err)
		}
	}
}
