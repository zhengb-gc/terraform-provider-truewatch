A workspace member is a user who has access to a workspace.

TrueWatch Cloud supports managing all members of the current workspace through member management, including setting role permissions, inviting members and setting permissions for members, configuring member groups, and setting SSO single sign-on.

Relationships:

```mermaid
graph LR
    A[Workspace] --> B[Member]
```

## Example Usage

```terraform
variable "email" {
  type = string
}

data "truewatch_members" "demo" {
  search = var.email
}

output "member" {
  value = data.truewatch_members.demo.members
}
```