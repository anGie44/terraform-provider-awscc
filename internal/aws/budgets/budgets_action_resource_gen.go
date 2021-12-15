// Code generated by generators/resource/main.go; DO NOT EDIT.

package budgets

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_budgets_budgets_action", budgetsActionResourceType)
}

// budgetsActionResourceType returns the Terraform awscc_budgets_budgets_action resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Budgets::BudgetsAction resource type.
func budgetsActionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"action_id": {
			// Property: ActionId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"action_threshold": {
			// Property: ActionThreshold
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Type": {
			//       "enum": [
			//         "PERCENTAGE",
			//         "ABSOLUTE_VALUE"
			//       ],
			//       "type": "string"
			//     },
			//     "Value": {
			//       "type": "number"
			//     }
			//   },
			//   "required": [
			//     "Value",
			//     "Type"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"type": {
						// Property: Type
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"PERCENTAGE",
								"ABSOLUTE_VALUE",
							}),
						},
					},
					"value": {
						// Property: Value
						Type:     types.NumberType,
						Required: true,
					},
				},
			),
			Required: true,
		},
		"action_type": {
			// Property: ActionType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "APPLY_IAM_POLICY",
			//     "APPLY_SCP_POLICY",
			//     "RUN_SSM_DOCUMENTS"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"APPLY_IAM_POLICY",
					"APPLY_SCP_POLICY",
					"RUN_SSM_DOCUMENTS",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"approval_model": {
			// Property: ApprovalModel
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "AUTOMATIC",
			//     "MANUAL"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"AUTOMATIC",
					"MANUAL",
				}),
			},
		},
		"budget_name": {
			// Property: BudgetName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"definition": {
			// Property: Definition
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "IamActionDefinition": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Groups": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "maxItems": 100,
			//           "minItems": 1,
			//           "type": "array"
			//         },
			//         "PolicyArn": {
			//           "type": "string"
			//         },
			//         "Roles": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "maxItems": 100,
			//           "minItems": 1,
			//           "type": "array"
			//         },
			//         "Users": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "maxItems": 100,
			//           "minItems": 1,
			//           "type": "array"
			//         }
			//       },
			//       "required": [
			//         "PolicyArn"
			//       ],
			//       "type": "object"
			//     },
			//     "ScpActionDefinition": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "PolicyId": {
			//           "type": "string"
			//         },
			//         "TargetIds": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "maxItems": 100,
			//           "minItems": 1,
			//           "type": "array"
			//         }
			//       },
			//       "required": [
			//         "PolicyId",
			//         "TargetIds"
			//       ],
			//       "type": "object"
			//     },
			//     "SsmActionDefinition": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "InstanceIds": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "maxItems": 100,
			//           "minItems": 1,
			//           "type": "array"
			//         },
			//         "Region": {
			//           "type": "string"
			//         },
			//         "Subtype": {
			//           "enum": [
			//             "STOP_EC2_INSTANCES",
			//             "STOP_RDS_INSTANCES"
			//           ],
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "Subtype",
			//         "Region",
			//         "InstanceIds"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"iam_action_definition": {
						// Property: IamActionDefinition
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"groups": {
									// Property: Groups
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenBetween(1, 100),
									},
								},
								"policy_arn": {
									// Property: PolicyArn
									Type:     types.StringType,
									Required: true,
								},
								"roles": {
									// Property: Roles
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenBetween(1, 100),
									},
								},
								"users": {
									// Property: Users
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenBetween(1, 100),
									},
								},
							},
						),
						Optional: true,
					},
					"scp_action_definition": {
						// Property: ScpActionDefinition
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"policy_id": {
									// Property: PolicyId
									Type:     types.StringType,
									Required: true,
								},
								"target_ids": {
									// Property: TargetIds
									Type:     types.ListType{ElemType: types.StringType},
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenBetween(1, 100),
									},
								},
							},
						),
						Optional: true,
					},
					"ssm_action_definition": {
						// Property: SsmActionDefinition
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"instance_ids": {
									// Property: InstanceIds
									Type:     types.ListType{ElemType: types.StringType},
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenBetween(1, 100),
									},
								},
								"region": {
									// Property: Region
									Type:     types.StringType,
									Required: true,
								},
								"subtype": {
									// Property: Subtype
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"STOP_EC2_INSTANCES",
											"STOP_RDS_INSTANCES",
										}),
									},
								},
							},
						),
						Optional: true,
					},
				},
			),
			Required: true,
		},
		"execution_role_arn": {
			// Property: ExecutionRoleArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
		},
		"notification_type": {
			// Property: NotificationType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "ACTUAL",
			//     "FORECASTED"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ACTUAL",
					"FORECASTED",
				}),
			},
		},
		"subscribers": {
			// Property: Subscribers
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Address": {
			//         "type": "string"
			//       },
			//       "Type": {
			//         "enum": [
			//           "SNS",
			//           "EMAIL"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Type",
			//       "Address"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 11,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"address": {
						// Property: Address
						Type:     types.StringType,
						Required: true,
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"SNS",
								"EMAIL",
							}),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(1, 11),
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
		Description: "An example resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Budgets::BudgetsAction").WithTerraformTypeName("awscc_budgets_budgets_action")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action_id":             "ActionId",
		"action_threshold":      "ActionThreshold",
		"action_type":           "ActionType",
		"address":               "Address",
		"approval_model":        "ApprovalModel",
		"budget_name":           "BudgetName",
		"definition":            "Definition",
		"execution_role_arn":    "ExecutionRoleArn",
		"groups":                "Groups",
		"iam_action_definition": "IamActionDefinition",
		"instance_ids":          "InstanceIds",
		"notification_type":     "NotificationType",
		"policy_arn":            "PolicyArn",
		"policy_id":             "PolicyId",
		"region":                "Region",
		"roles":                 "Roles",
		"scp_action_definition": "ScpActionDefinition",
		"ssm_action_definition": "SsmActionDefinition",
		"subscribers":           "Subscribers",
		"subtype":               "Subtype",
		"target_ids":            "TargetIds",
		"type":                  "Type",
		"users":                 "Users",
		"value":                 "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
