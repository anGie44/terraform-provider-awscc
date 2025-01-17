---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_budgets_budgets_action Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  An example resource schema demonstrating some basic constructs and validation rules.
---

# awscc_budgets_budgets_action (Resource)

An example resource schema demonstrating some basic constructs and validation rules.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `action_threshold` (Attributes) (see [below for nested schema](#nestedatt--action_threshold))
- `action_type` (String)
- `budget_name` (String)
- `definition` (Attributes) (see [below for nested schema](#nestedatt--definition))
- `execution_role_arn` (String)
- `notification_type` (String)
- `subscribers` (Attributes List) (see [below for nested schema](#nestedatt--subscribers))

### Optional

- `approval_model` (String)

### Read-Only

- `action_id` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--action_threshold"></a>
### Nested Schema for `action_threshold`

Required:

- `type` (String)
- `value` (Number)


<a id="nestedatt--definition"></a>
### Nested Schema for `definition`

Required:

- `iam_action_definition` (Attributes) (see [below for nested schema](#nestedatt--definition--iam_action_definition))
- `scp_action_definition` (Attributes) (see [below for nested schema](#nestedatt--definition--scp_action_definition))
- `ssm_action_definition` (Attributes) (see [below for nested schema](#nestedatt--definition--ssm_action_definition))

<a id="nestedatt--definition--iam_action_definition"></a>
### Nested Schema for `definition.iam_action_definition`

Required:

- `groups` (List of String)
- `policy_arn` (String)
- `roles` (List of String)
- `users` (List of String)


<a id="nestedatt--definition--scp_action_definition"></a>
### Nested Schema for `definition.scp_action_definition`

Required:

- `policy_id` (String)
- `target_ids` (List of String)


<a id="nestedatt--definition--ssm_action_definition"></a>
### Nested Schema for `definition.ssm_action_definition`

Required:

- `instance_ids` (List of String)
- `region` (String)
- `subtype` (String)



<a id="nestedatt--subscribers"></a>
### Nested Schema for `subscribers`

Required:

- `address` (String)
- `type` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_budgets_budgets_action.example <resource ID>
```
