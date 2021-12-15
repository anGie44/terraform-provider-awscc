// Code generated by generators/resource/main.go; DO NOT EDIT.

package resiliencehub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_resiliencehub_app", appResourceType)
}

// appResourceType returns the Terraform awscc_resiliencehub_app resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ResilienceHub::App resource type.
func appResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"app_arn": {
			// Property: AppArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name (ARN) of the App.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Amazon Resource Name (ARN) of the App.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"app_template_body": {
			// Property: AppTemplateBody
			// CloudFormation resource type schema:
			// {
			//   "description": "A string containing full ResilienceHub app template body.",
			//   "maxLength": 5000,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A string containing full ResilienceHub app template body.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 5000),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "App description.",
			//   "maxLength": 500,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "App description.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 500),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the app.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Name of the app.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"resiliency_policy_arn": {
			// Property: ResiliencyPolicyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name (ARN) of the Resiliency Policy.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Amazon Resource Name (ARN) of the Resiliency Policy.",
			Type:        types.StringType,
			Optional:    true,
		},
		"resource_mappings": {
			// Property: ResourceMappings
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of ResourceMapping objects.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Resource mapping is used to map logical resources from template to physical resource",
			//     "properties": {
			//       "LogicalStackName": {
			//         "type": "string"
			//       },
			//       "MappingType": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "PhysicalResourceId": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "AwsAccountId": {
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "AwsRegion": {
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "Identifier": {
			//             "maxLength": 255,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "Type": {
			//             "pattern": "",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Identifier",
			//           "Type"
			//         ],
			//         "type": "object"
			//       },
			//       "ResourceName": {
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "MappingType",
			//       "PhysicalResourceId"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "An array of ResourceMapping objects.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"logical_stack_name": {
						// Property: LogicalStackName
						Type:     types.StringType,
						Optional: true,
					},
					"mapping_type": {
						// Property: MappingType
						Type:     types.StringType,
						Required: true,
					},
					"physical_resource_id": {
						// Property: PhysicalResourceId
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"aws_account_id": {
									// Property: AwsAccountId
									Type:     types.StringType,
									Optional: true,
								},
								"aws_region": {
									// Property: AwsRegion
									Type:     types.StringType,
									Optional: true,
								},
								"identifier": {
									// Property: Identifier
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 255),
									},
								},
								"type": {
									// Property: Type
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Required: true,
					},
					"resource_name": {
						// Property: ResourceName
						Type:     types.StringType,
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "patternProperties": {
			//     "": {
			//       "maxLength": 256,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
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
		Description: "Resource Type Definition for AWS::ResilienceHub::App.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ResilienceHub::App").WithTerraformTypeName("awscc_resiliencehub_app")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_arn":               "AppArn",
		"app_template_body":     "AppTemplateBody",
		"aws_account_id":        "AwsAccountId",
		"aws_region":            "AwsRegion",
		"description":           "Description",
		"identifier":            "Identifier",
		"logical_stack_name":    "LogicalStackName",
		"mapping_type":          "MappingType",
		"name":                  "Name",
		"physical_resource_id":  "PhysicalResourceId",
		"resiliency_policy_arn": "ResiliencyPolicyArn",
		"resource_mappings":     "ResourceMappings",
		"resource_name":         "ResourceName",
		"tags":                  "Tags",
		"type":                  "Type",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
