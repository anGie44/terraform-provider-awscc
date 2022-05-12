// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package logs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_logs_resource_policy", resourcePolicyDataSourceType)
}

// resourcePolicyDataSourceType returns the Terraform awscc_logs_resource_policy data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Logs::ResourcePolicy resource type.
func resourcePolicyDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"policy_document": {
			// Property: PolicyDocument
			// CloudFormation resource type schema:
			// {
			//   "description": "The policy document",
			//   "maxLength": 5120,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The policy document",
			Type:        types.StringType,
			Computed:    true,
		},
		"policy_name": {
			// Property: PolicyName
			// CloudFormation resource type schema:
			// {
			//   "description": "A name for resource policy",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "^([^:*\\/]+\\/?)*[^:*\\/]+$",
			//   "type": "string"
			// }
			Description: "A name for resource policy",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Logs::ResourcePolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::ResourcePolicy").WithTerraformTypeName("awscc_logs_resource_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"policy_document": "PolicyDocument",
		"policy_name":     "PolicyName",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
