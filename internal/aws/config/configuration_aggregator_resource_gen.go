// Code generated by generators/resource/main.go; DO NOT EDIT.

package config

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
	registry.AddResourceTypeFactory("awscc_config_configuration_aggregator", configurationAggregatorResourceType)
}

// configurationAggregatorResourceType returns the Terraform awscc_config_configuration_aggregator resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Config::ConfigurationAggregator resource type.
func configurationAggregatorResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"account_aggregation_sources": {
			// Property: AccountAggregationSources
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "AccountIds": {
			//         "items": {
			//           "type": "string"
			//         },
			//         "type": "array",
			//         "uniqueItems": false
			//       },
			//       "AllAwsRegions": {
			//         "type": "boolean"
			//       },
			//       "AwsRegions": {
			//         "items": {
			//           "type": "string"
			//         },
			//         "type": "array",
			//         "uniqueItems": false
			//       }
			//     },
			//     "required": [
			//       "AccountIds"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"account_ids": {
						// Property: AccountIds
						Type:     types.ListType{ElemType: types.StringType},
						Required: true,
					},
					"all_aws_regions": {
						// Property: AllAwsRegions
						Type:     types.BoolType,
						Optional: true,
					},
					"aws_regions": {
						// Property: AwsRegions
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"configuration_aggregator_arn": {
			// Property: ConfigurationAggregatorArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the aggregator.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the aggregator.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"configuration_aggregator_name": {
			// Property: ConfigurationAggregatorName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the aggregator.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "[\\w\\-]+",
			//   "type": "string"
			// }
			Description: "The name of the aggregator.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
				validate.StringMatch(regexp.MustCompile("[\\w\\-]+"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"organization_aggregation_source": {
			// Property: OrganizationAggregationSource
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AllAwsRegions": {
			//       "type": "boolean"
			//     },
			//     "AwsRegions": {
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "RoleArn": {
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "RoleArn"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"all_aws_regions": {
						// Property: AllAwsRegions
						Type:     types.BoolType,
						Optional: true,
					},
					"aws_regions": {
						// Property: AwsRegions
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
					},
					"role_arn": {
						// Property: RoleArn
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Optional: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags for the configuration aggregator.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The tags for the configuration aggregator.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
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
		Description: "Resource Type definition for AWS::Config::ConfigurationAggregator",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Config::ConfigurationAggregator").WithTerraformTypeName("awscc_config_configuration_aggregator")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_aggregation_sources":     "AccountAggregationSources",
		"account_ids":                     "AccountIds",
		"all_aws_regions":                 "AllAwsRegions",
		"aws_regions":                     "AwsRegions",
		"configuration_aggregator_arn":    "ConfigurationAggregatorArn",
		"configuration_aggregator_name":   "ConfigurationAggregatorName",
		"key":                             "Key",
		"organization_aggregation_source": "OrganizationAggregationSource",
		"role_arn":                        "RoleArn",
		"tags":                            "Tags",
		"value":                           "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
