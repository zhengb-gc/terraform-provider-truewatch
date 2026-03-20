package Monitor_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/provider"
)

func TestAccMonitor(t *testing.T) {
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

resource "truewatch_alertpolicy" "demo" {
  name           = "oac-demo"
  silent_timeout = "1h"

  statuses = [
    "critical",
    "error",
    "warning",
    "info",
    "ok",
    "nodata",
    "nodata_ok",
    "nodata_as_ok",
  ]

  alert_targets = [
    {
      type         = "member_group"
      member_group = {
        id = truewatch_membergroup.demo.id
      }
    },
  ]
}

resource "truewatch_monitor" "demo" {
  manifest     = file("${path.module}/monitor.json")
  alert_policy = {
    id = truewatch_alertpolicy.demo.id
  }
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}
