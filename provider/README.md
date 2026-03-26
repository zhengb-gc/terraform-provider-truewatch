# Terraform Provider: TrueWatch

The TrueWatch Provider provides resources to manage [TrueWatch Cloud](https://en.truewatch.com/) resources.

To learn the basics of Terraform using this provider, follow the hands-on get started tutorials.

Interested in the provider's latest features, or want to make sure you're up to date? Check out the changelog for version information and release notes.

## Authenticating to TrueWatch Cloud

Terraform supports a number of different methods for authenticating to TrueWatch Cloud:

* [Workspace Key](https://console.truewatch.com/workspace/apiManage)

## Usage

```terraform
# We strongly recommend using the required_providers block to set the
# TrueWatch Cloud Provider source and version being used
terraform {
  required_version = ">=0.12"

  required_providers {
    truewatch = {
      source = "TrueWatchTech/truewatch"
      version = "=0.0.2"
    }
  }
}

// We also recommend use secret environment variables to set the provider,
// Such as TRUEWATCH_ACCESS_TOKEN and TRUEWATCH_REGION
provider "truewatch" {
  # access_token = "your access token, recommend store in environment variable"
  region = "singapore"
  # end_point = "https://openapi.truewatch.com"
}
```

## More Examples

* [Example Source Code](https://github.com/TrueWatchTech/terraform-provider-truewatch/tree/main/examples)
