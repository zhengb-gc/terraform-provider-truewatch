package Membergroup_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/provider"
)

func TestAccMembergroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: provider.Config + `
variable "email" {
  type = string
}

data "truewatch_members" "demo" {
  filters = [
    {
      name   = "email"
      values = [var.email]
    }
  ]
}

resource "truewatch_membergroup" "demo" {
  name       = "oac-demo"
  member_ids = data.truewatch_members.demo.items[*].id
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}
