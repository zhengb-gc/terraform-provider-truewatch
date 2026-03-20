A role permission is defined to grant a user permission to perform a specific action on a specific resource.

TrueWatch Cloud supports a number of different role permissions, and you can create new roles for users and assign permission scopes to roles to meet the permission needs of your organization.


## Example Usage

```terraform
data "truewatch_permissions" "demo" {
}

output "permissions" {
  value = data.truewatch_permissions.demo
}
```