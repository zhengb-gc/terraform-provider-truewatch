variable "email" {
  type = string
}

data "truewatch_members" "demo" {
  search = var.email
}

resource "truewatch_membergroup" "demo" {
  name          = "oac-demo2"
  account_uuids = data.truewatch_members.demo.members[*].uuid
}

output "member" {
  value = data.truewatch_members.demo.members
}
