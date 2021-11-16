// Code generated by generators/resource/main.go; DO NOT EDIT.

package connect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_connect_user_hierarchy_group", userHierarchyGroupResourceType)
}

// userHierarchyGroupResourceType returns the Terraform awscc_connect_user_hierarchy_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Connect::UserHierarchyGroup resource type.
func userHierarchyGroupResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"instance_arn": {
			// Property: InstanceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the Amazon Connect instance.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The identifier of the Amazon Connect instance.",
			Type:        types.StringType,
			Required:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the user hierarchy group.",
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the user hierarchy group.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 100),
			},
		},
		"parent_group_arn": {
			// Property: ParentGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) for the parent user hierarchy group.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) for the parent user hierarchy group.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"user_hierarchy_group_arn": {
			// Property: UserHierarchyGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) for the user hierarchy group.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) for the user hierarchy group.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::Connect::UserHierarchyGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::UserHierarchyGroup").WithTerraformTypeName("awscc_connect_user_hierarchy_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"instance_arn":             "InstanceArn",
		"name":                     "Name",
		"parent_group_arn":         "ParentGroupArn",
		"user_hierarchy_group_arn": "UserHierarchyGroupArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
