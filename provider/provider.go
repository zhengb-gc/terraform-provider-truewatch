package provider

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/api"
)

//go:embed README.md
var doc string

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &truewatchProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &truewatchProvider{}
}

// truewatchProvider is the provider implementation.
type truewatchProvider struct{}

// truewatchProviderModel maps provider schema data to a Go type.
type truewatchProviderModel struct {
	AccessToken types.String `tfsdk:"access_token"`
	EndPoint    types.String `tfsdk:"end_point"`
	Region      types.String `tfsdk:"region"`
}

// Metadata returns the provider type name.
func (p *truewatchProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "truewatch"
}

// Schema defines the provider-level schema for configuration data.
func (p *truewatchProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Interact with TrueWatch Cloud.",
		MarkdownDescription: doc,
		Attributes: map[string]schema.Attribute{
			"region": schema.StringAttribute{
				Description:         "Region for TrueWatch Cloud API. May also be provided via TRUEWATCH_REGION environment variable. See https://github.com/TrueWatchTech/terraform-provider-truewatch for a list of available regions.",
				MarkdownDescription: "Region for TrueWatch Cloud API. May also be provided via TRUEWATCH_REGION environment variable. See [GitHub](https://github.com/TrueWatchTech/terraform-provider-truewatch) for a list of available regions.",
				Optional:            true,
			},
			"end_point": schema.StringAttribute{
				Description:         "EndPoint for TrueWatch Cloud API. May also be provided via TRUEWATCH_END_POINT environment variable. See https://github.com/TrueWatchTech/terraform-provider-truewatch for a list of available regions.",
				MarkdownDescription: "EndPoint for TrueWatch Cloud API. May also be provided via TRUEWATCH_END_POINT environment variable. See [GitHub](https://github.com/TrueWatchTech/terraform-provider-truewatch) for a list of available regions.",
				Optional:            true,
			},
			"access_token": schema.StringAttribute{
				Description:         "Access token for TrueWatch Cloud API. May also be provided via TRUEWATCH_ACCESS_TOKEN environment variable. Get an Key ID from https://console.truewatch.com/workspace/apiManage as access token.",
				MarkdownDescription: "Access token for TrueWatch Cloud API. May also be provided via TRUEWATCH_ACCESS_TOKEN environment variable. Get an Key ID from [TrueWatch Cloud](https://console.truewatch.com/workspace/apiManage) as access token.",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

// Configure prepares a TrueWatch Cloud API client for data sources and resources.
func (p *truewatchProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config truewatchProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	accessToken := getConfigField("access_token", config.AccessToken, false, resp)
	if resp.Diagnostics.HasError() {
		return
	}

	region := getConfigField("region", config.Region, true, resp)
	endPoint := getConfigField("end_point", config.EndPoint, true, resp)

	ctx = tflog.SetField(ctx, "truewatch_region", region)
	ctx = tflog.SetField(ctx, "truewatch_end_point", endPoint)
	ctx = tflog.SetField(ctx, "truewatch_access_token", accessToken)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "truewatch_access_token")

	tflog.Debug(ctx, "Creating TrueWatch Cloud client")

	// Create a new TrueWatch Cloud client using the configuration values
	client, err := api.NewClient(region, accessToken, endPoint)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create TrueWatch Cloud API Client",
			"An unexpected error occurred when creating the TrueWatch Cloud API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"TrueWatch Cloud Client Error: "+err.Error(),
		)
		return
	}

	// Make the TrueWatch Cloud client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
}

func getConfigField(name string, value types.String, allowEmpty bool, resp *provider.ConfigureResponse) string {
	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	envName := fmt.Sprintf("TRUEWATCH_%s", strings.ToUpper(name))

	if value.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root(name),
			fmt.Sprintf("Unknown TrueWatch Cloud API %s", strings.ToTitle(name)),
			"The provider cannot create the TrueWatch Cloud API client as there is an unknown configuration value for the TrueWatch Cloud API endpoint. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the "+envName+" environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return ""
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	valueString := os.Getenv(envName)
	if !value.IsNull() {
		valueString = value.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if !allowEmpty && valueString == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root(name),
			fmt.Sprintf("Missing TrueWatch Cloud API %s", strings.ToTitle(name)),
			"The provider cannot create the TrueWatch Cloud API client as there is a missing or empty value for the TrueWatch Cloud API "+name+". "+
				"Set the host value in the configuration or use the "+envName+" environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	return valueString
}
