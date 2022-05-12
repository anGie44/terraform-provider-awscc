// Code generated by generators/resource/main.go; DO NOT EDIT.

package backup

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
	registry.AddResourceTypeFactory("awscc_backup_framework", frameworkResourceType)
}

// frameworkResourceType returns the Terraform awscc_backup_framework resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Backup::Framework resource type.
func frameworkResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The date and time that a framework is created, in Unix format and Coordinated Universal Time (UTC). The value of `CreationTime` is accurate to milliseconds. For example, the value 1516925490.087 represents Friday, January 26, 2018 12:11:30.087 AM.",
			//   "type": "number"
			// }
			Description: "The date and time that a framework is created, in Unix format and Coordinated Universal Time (UTC). The value of `CreationTime` is accurate to milliseconds. For example, the value 1516925490.087 represents Friday, January 26, 2018 12:11:30.087 AM.",
			Type:        types.Float64Type,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"deployment_status": {
			// Property: DeploymentStatus
			// CloudFormation resource type schema:
			// {
			//   "description": "The deployment status of a framework. The statuses are: `CREATE_IN_PROGRESS | UPDATE_IN_PROGRESS | DELETE_IN_PROGRESS | COMPLETED | FAILED`",
			//   "type": "string"
			// }
			Description: "The deployment status of a framework. The statuses are: `CREATE_IN_PROGRESS | UPDATE_IN_PROGRESS | DELETE_IN_PROGRESS | COMPLETED | FAILED`",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"framework_arn": {
			// Property: FrameworkArn
			// CloudFormation resource type schema:
			// {
			//   "description": "An Amazon Resource Name (ARN) that uniquely identifies Framework as a resource",
			//   "type": "string"
			// }
			Description: "An Amazon Resource Name (ARN) that uniquely identifies Framework as a resource",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"framework_controls": {
			// Property: FrameworkControls
			// CloudFormation resource type schema:
			// {
			//   "description": "Contains detailed information about all of the controls of a framework. Each framework must contain at least one control.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ControlInputParameters": {
			//         "description": "A list of ParameterName and ParameterValue pairs.",
			//         "insertionOrder": false,
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "ParameterName": {
			//               "type": "string"
			//             },
			//             "ParameterValue": {
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "ParameterName",
			//             "ParameterValue"
			//           ],
			//           "type": "object"
			//         },
			//         "type": "array",
			//         "uniqueItems": true
			//       },
			//       "ControlName": {
			//         "description": "The name of a control. This name is between 1 and 256 characters.",
			//         "type": "string"
			//       },
			//       "ControlScope": {
			//         "additionalProperties": false,
			//         "description": "The scope of a control. The control scope defines what the control will evaluate. Three examples of control scopes are: a specific backup plan, all backup plans with a specific tag, or all backup plans.",
			//         "properties": {
			//           "ComplianceResourceIds": {
			//             "description": "The ID of the only AWS resource that you want your control scope to contain.",
			//             "insertionOrder": false,
			//             "items": {
			//               "type": "string"
			//             },
			//             "type": "array"
			//           },
			//           "ComplianceResourceTypes": {
			//             "description": "Describes whether the control scope includes one or more types of resources, such as `EFS` or `RDS`.",
			//             "insertionOrder": false,
			//             "items": {
			//               "type": "string"
			//             },
			//             "type": "array"
			//           },
			//           "Tags": {
			//             "description": "Describes whether the control scope includes resources with one or more tags. Each tag is a key-value pair.",
			//             "insertionOrder": false,
			//             "items": {
			//               "additionalProperties": false,
			//               "description": "A key-value pair to associate with a resource.",
			//               "properties": {
			//                 "Key": {
			//                   "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//                   "maxLength": 128,
			//                   "minLength": 1,
			//                   "type": "string"
			//                 },
			//                 "Value": {
			//                   "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//                   "maxLength": 256,
			//                   "minLength": 0,
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Key",
			//                 "Value"
			//               ],
			//               "type": "object"
			//             },
			//             "type": "array"
			//           }
			//         },
			//         "type": "object"
			//       }
			//     },
			//     "required": [
			//       "ControlName"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Contains detailed information about all of the controls of a framework. Each framework must contain at least one control.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"control_input_parameters": {
						// Property: ControlInputParameters
						Description: "A list of ParameterName and ParameterValue pairs.",
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"parameter_name": {
									// Property: ParameterName
									Type:     types.StringType,
									Required: true,
								},
								"parameter_value": {
									// Property: ParameterValue
									Type:     types.StringType,
									Required: true,
								},
							},
							tfsdk.SetNestedAttributesOptions{},
						),
						Optional: true,
					},
					"control_name": {
						// Property: ControlName
						Description: "The name of a control. This name is between 1 and 256 characters.",
						Type:        types.StringType,
						Required:    true,
					},
					"control_scope": {
						// Property: ControlScope
						Description: "The scope of a control. The control scope defines what the control will evaluate. Three examples of control scopes are: a specific backup plan, all backup plans with a specific tag, or all backup plans.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"compliance_resource_ids": {
									// Property: ComplianceResourceIds
									Description: "The ID of the only AWS resource that you want your control scope to contain.",
									Type:        types.ListType{ElemType: types.StringType},
									Optional:    true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										Multiset(),
									},
								},
								"compliance_resource_types": {
									// Property: ComplianceResourceTypes
									Description: "Describes whether the control scope includes one or more types of resources, such as `EFS` or `RDS`.",
									Type:        types.ListType{ElemType: types.StringType},
									Optional:    true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										Multiset(),
									},
								},
								"tags": {
									// Property: Tags
									Description: "Describes whether the control scope includes resources with one or more tags. Each tag is a key-value pair.",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"key": {
												// Property: Key
												Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
												Type:        types.StringType,
												Required:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 128),
												},
											},
											"value": {
												// Property: Value
												Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
									PlanModifiers: []tfsdk.AttributePlanModifier{
										Multiset(),
									},
								},
							},
						),
						Optional: true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Required: true,
		},
		"framework_description": {
			// Property: FrameworkDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "An optional description of the framework with a maximum 1,024 characters.",
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "An optional description of the framework with a maximum 1,024 characters.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 1024),
			},
		},
		"framework_name": {
			// Property: FrameworkName
			// CloudFormation resource type schema:
			// {
			//   "description": "The unique name of a framework. This name is between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "[a-zA-Z][_a-zA-Z0-9]*",
			//   "type": "string"
			// }
			Description: "The unique name of a framework. This name is between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
				validate.StringMatch(regexp.MustCompile("[a-zA-Z][_a-zA-Z0-9]*"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"framework_status": {
			// Property: FrameworkStatus
			// CloudFormation resource type schema:
			// {
			//   "description": "A framework consists of one or more controls. Each control governs a resource, such as backup plans, backup selections, backup vaults, or recovery points. You can also turn AWS Config recording on or off for each resource. The statuses are:\n\n`ACTIVE` when recording is turned on for all resources governed by the framework.\n\n`PARTIALLY_ACTIVE` when recording is turned off for at least one resource governed by the framework.\n\n`INACTIVE` when recording is turned off for all resources governed by the framework.\n\n`UNAVAILABLE` when AWS Backup is unable to validate recording status at this time.",
			//   "type": "string"
			// }
			Description: "A framework consists of one or more controls. Each control governs a resource, such as backup plans, backup selections, backup vaults, or recovery points. You can also turn AWS Config recording on or off for each resource. The statuses are:\n\n`ACTIVE` when recording is turned on for all resources governed by the framework.\n\n`PARTIALLY_ACTIVE` when recording is turned off for at least one resource governed by the framework.\n\n`INACTIVE` when recording is turned off for all resources governed by the framework.\n\n`UNAVAILABLE` when AWS Backup is unable to validate recording status at this time.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"framework_tags": {
			// Property: FrameworkTags
			// CloudFormation resource type schema:
			// {
			//   "description": "Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
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
		Description: "Contains detailed information about a framework. Frameworks contain controls, which evaluate and report on your backup events and resources. Frameworks generate daily compliance results.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Backup::Framework").WithTerraformTypeName("awscc_backup_framework")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"compliance_resource_ids":   "ComplianceResourceIds",
		"compliance_resource_types": "ComplianceResourceTypes",
		"control_input_parameters":  "ControlInputParameters",
		"control_name":              "ControlName",
		"control_scope":             "ControlScope",
		"creation_time":             "CreationTime",
		"deployment_status":         "DeploymentStatus",
		"framework_arn":             "FrameworkArn",
		"framework_controls":        "FrameworkControls",
		"framework_description":     "FrameworkDescription",
		"framework_name":            "FrameworkName",
		"framework_status":          "FrameworkStatus",
		"framework_tags":            "FrameworkTags",
		"key":                       "Key",
		"parameter_name":            "ParameterName",
		"parameter_value":           "ParameterValue",
		"tags":                      "Tags",
		"value":                     "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
