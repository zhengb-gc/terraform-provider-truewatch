package Blacklist_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/provider"
)

func TestAccBlacklist(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: provider.Config + `
resource "truewatch_blacklist" "demo" {
  name = "test-blacklist"
  type = "logging"
  desc = "Test blacklist"
  source = "nginx"

  filters = [
    {
      name      = "foo"
      operation = "in"
      condition = "and"
      values    = ["oac-*"]
    }
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("truewatch_blacklist.demo", "name", "test-blacklist"),
					resource.TestCheckResourceAttr("truewatch_blacklist.demo", "type", "logging"),
					resource.TestCheckResourceAttr("truewatch_blacklist.demo", "desc", "Test blacklist"),
					resource.TestCheckResourceAttr("truewatch_blacklist.demo", "source", "nginx"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}
