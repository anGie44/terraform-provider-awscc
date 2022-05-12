// Code generated by generators/resource/main.go; DO NOT EDIT.

package elasticache

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_elasticache_user_group", userGroupResourceType)
}

// userGroupResourceType returns the Terraform awscc_elasticache_user_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ElastiCache::UserGroup resource type.
func userGroupResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the user account.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the user account.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"engine": {
			// Property: Engine
			// CloudFormation resource type schema:
			// {
			//   "description": "Must be redis.",
			//   "enum": [
			//     "redis"
			//   ],
			//   "type": "string"
			// }
			Description: "Must be redis.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"redis",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates user group status. Can be \"creating\", \"active\", \"modifying\", \"deleting\".",
			//   "type": "string"
			// }
			Description: "Indicates user group status. Can be \"creating\", \"active\", \"modifying\", \"deleting\".",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"user_group_id": {
			// Property: UserGroupId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the user group.",
			//   "pattern": "[a-z][a-z0-9\\\\-]*",
			//   "type": "string"
			// }
			Description: "The ID of the user group.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("[a-z][a-z0-9\\\\-]*"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"user_ids": {
			// Property: UserIds
			// CloudFormation resource type schema:
			// {
			//   "$comment": "List of users.",
			//   "description": "List of users associated to this user group.",
			//   "insertionOrder": true,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "List of users associated to this user group.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::ElastiCache::UserGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElastiCache::UserGroup").WithTerraformTypeName("awscc_elasticache_user_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":           "Arn",
		"engine":        "Engine",
		"status":        "Status",
		"user_group_id": "UserGroupId",
		"user_ids":      "UserIds",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
