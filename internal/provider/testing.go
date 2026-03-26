package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"

	"github.com/TrueWatchTech/terraform-provider-truewatch/provider"
)

const (
	// Config is a shared configuration to combine with the actual
	// test configuration so the TrueWatch Cloud client is properly configured.
	// It is also possible to use the TRUEWATCH_ environment variables instead,
	// such as updating the Makefile and running the testing through that tool.
	Config = `
terraform {
	required_version = ">=0.12"

	required_providers {
		truewatch = {
			source = "TrueWatchTech/truewatch"
		}
	}
}

provider "truewatch" {
	region = "singapore"
	token = ""
}
`
)

var (
	// TestAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	TestAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"truewatch": providerserver.NewProtocol6WithError(provider.New()),
	}
)
