---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_resiliencehub_resiliency_policy Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type Definition for Resiliency Policy.
---

# awscc_resiliencehub_resiliency_policy (Resource)

Resource Type Definition for Resiliency Policy.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `policy` (Attributes Map) (see [below for nested schema](#nestedatt--policy))
- `policy_name` (String) Name of Resiliency Policy.
- `tier` (String) Resiliency Policy Tier.

### Optional

- `data_location_constraint` (String) Data Location Constraint of the Policy.
- `policy_description` (String) Description of Resiliency Policy.
- `tags` (Map of String)

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `policy_arn` (String) Amazon Resource Name (ARN) of the Resiliency Policy.

<a id="nestedatt--policy"></a>
### Nested Schema for `policy`

Required:

- `rpo_in_secs` (Number) RPO in seconds.
- `rto_in_secs` (Number) RTO in seconds.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_resiliencehub_resiliency_policy.example <resource ID>
```
