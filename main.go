package main

import (
	"./katakunai"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: katakunai.Provider,
	})
}
