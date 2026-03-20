resource "truewatch_role" "role" {
  name = "tf-test-role1"
  desc = "test role"
  keys = ["snapshot.delete", "workspace.readMember"]
}
