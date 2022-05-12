---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_route53resolver_resolver_rule_association Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Route53Resolver::ResolverRuleAssociation
---

# awscc_route53resolver_resolver_rule_association (Data Source)

Data Source schema for AWS::Route53Resolver::ResolverRuleAssociation



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `name` (String) The name of an association between a Resolver rule and a VPC.
- `resolver_rule_association_id` (String) Primary Identifier for Resolver Rule Association
- `resolver_rule_id` (String) The ID of the Resolver rule that you associated with the VPC that is specified by VPCId.
- `vpc_id` (String) The ID of the VPC that you associated the Resolver rule with.

