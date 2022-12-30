package main

import providers "github.com/EnergoStalin/vsc-app/pkg/vsc-providers"

func main() {
	for _, provider := range providers.Providers {
		print(provider.GetName())
	}
}
