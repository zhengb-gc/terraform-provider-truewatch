package provider

import (
	"context"

	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/datasources/members"
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/datasources/permissions"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// DataSources defines the data sources implemented in the provider.
func (p *truewatchProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		members.NewMembersDataSource,
		permissions.NewPermissionsDataSource,
		// default_region.NewDefaultRegionDataSource,
	}
}
