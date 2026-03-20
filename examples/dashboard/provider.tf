terraform {
  required_providers {
    truewatch = {
      source = "TrueWatchTech/truewatch"
    }
  }
}

provider "truewatch" {
  # You can set your API key here or use the TRUEWATCH_ACCESS_TOKEN environment variable
  # access_token = "your-api-key"
}
