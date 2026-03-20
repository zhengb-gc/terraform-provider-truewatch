
terraform {
  required_version = ">=0.12"

  required_providers {
    truewatch = {
      source = "TrueWatchTech/truewatch"
    }
  }
}

provider "truewatch" {
  region = "hangzhou"
}
