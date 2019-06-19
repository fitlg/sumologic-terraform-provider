package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/fitlg/sumologic-terraform-provider/sumologic"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: sumologic.Provider,
	})
}
