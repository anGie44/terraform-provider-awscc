// Code generated by generators/resource/main.go; DO NOT EDIT.

package s3outposts

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
	registry.AddResourceTypeFactory("awscc_s3outposts_bucket_policy", bucketPolicyResourceType)
}

// bucketPolicyResourceType returns the Terraform awscc_s3outposts_bucket_policy resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::S3Outposts::BucketPolicy resource type.
func bucketPolicyResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"bucket": {
			// Property: Bucket
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the specified bucket.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "^arn:[^:]+:s3-outposts:[a-zA-Z0-9\\-]+:\\d{12}:outpost\\/[^:]+\\/bucket\\/[^:]+$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the specified bucket.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 2048),
				validate.StringMatch(regexp.MustCompile("^arn:[^:]+:s3-outposts:[a-zA-Z0-9\\-]+:\\d{12}:outpost\\/[^:]+\\/bucket\\/[^:]+$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"policy_document": {
			// Property: PolicyDocument
			// CloudFormation resource type schema:
			// {
			//   "description": "A policy document containing permissions to add to the specified bucket.",
			//   "type": "object"
			// }
			Description: "A policy document containing permissions to add to the specified bucket.",
			Type:        types.MapType{ElemType: types.StringType},
			Required:    true,
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
		Description: "Resource Type Definition for AWS::S3Outposts::BucketPolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3Outposts::BucketPolicy").WithTerraformTypeName("awscc_s3outposts_bucket_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bucket":          "Bucket",
		"policy_document": "PolicyDocument",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
