// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_iot_mitigation_action", mitigationActionDataSourceType)
}

// mitigationActionDataSourceType returns the Terraform awscc_iot_mitigation_action data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::IoT::MitigationAction resource type.
func mitigationActionDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"action_name": {
			// Property: ActionName
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique identifier for the mitigation action.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "[a-zA-Z0-9:_-]+",
			//   "type": "string"
			// }
			Description: "A unique identifier for the mitigation action.",
			Type:        types.StringType,
			Computed:    true,
		},
		"action_params": {
			// Property: ActionParams
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The set of parameters for this mitigation action. You can specify only one type of parameter (in other words, you can apply only one action for each defined mitigation action).",
			//   "properties": {
			//     "AddThingsToThingGroupParams": {
			//       "additionalProperties": false,
			//       "description": "Parameters to define a mitigation action that moves devices associated with a certificate to one or more specified thing groups, typically for quarantine.",
			//       "properties": {
			//         "OverrideDynamicGroups": {
			//           "description": "Specifies if this mitigation action can move the things that triggered the mitigation action out of one or more dynamic thing groups.",
			//           "type": "boolean"
			//         },
			//         "ThingGroupNames": {
			//           "description": "The list of groups to which you want to add the things that triggered the mitigation action.",
			//           "insertionOrder": false,
			//           "items": {
			//             "maxLength": 128,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "maxItems": 10,
			//           "minItems": 1,
			//           "type": "array",
			//           "uniqueItems": true
			//         }
			//       },
			//       "required": [
			//         "ThingGroupNames"
			//       ],
			//       "type": "object"
			//     },
			//     "EnableIoTLoggingParams": {
			//       "additionalProperties": false,
			//       "description": "Parameters to define a mitigation action that enables AWS IoT logging at a specified level of detail.",
			//       "properties": {
			//         "LogLevel": {
			//           "description": " Specifies which types of information are logged.",
			//           "enum": [
			//             "DEBUG",
			//             "INFO",
			//             "ERROR",
			//             "WARN"
			//           ],
			//           "type": "string"
			//         },
			//         "RoleArnForLogging": {
			//           "description": " The ARN of the IAM role used for logging.",
			//           "maxLength": 2048,
			//           "minLength": 20,
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "LogLevel",
			//         "RoleArnForLogging"
			//       ],
			//       "type": "object"
			//     },
			//     "PublishFindingToSnsParams": {
			//       "additionalProperties": false,
			//       "description": "Parameters, to define a mitigation action that publishes findings to Amazon SNS. You can implement your own custom actions in response to the Amazon SNS messages.",
			//       "properties": {
			//         "TopicArn": {
			//           "description": "The ARN of the topic to which you want to publish the findings.",
			//           "maxLength": 2048,
			//           "minLength": 20,
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "TopicArn"
			//       ],
			//       "type": "object"
			//     },
			//     "ReplaceDefaultPolicyVersionParams": {
			//       "additionalProperties": false,
			//       "description": "Parameters to define a mitigation action that adds a blank policy to restrict permissions.",
			//       "properties": {
			//         "TemplateName": {
			//           "enum": [
			//             "BLANK_POLICY"
			//           ],
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "TemplateName"
			//       ],
			//       "type": "object"
			//     },
			//     "UpdateCACertificateParams": {
			//       "additionalProperties": false,
			//       "description": "Parameters to define a mitigation action that changes the state of the CA certificate to inactive.",
			//       "properties": {
			//         "Action": {
			//           "enum": [
			//             "DEACTIVATE"
			//           ],
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "Action"
			//       ],
			//       "type": "object"
			//     },
			//     "UpdateDeviceCertificateParams": {
			//       "additionalProperties": false,
			//       "description": "Parameters to define a mitigation action that changes the state of the device certificate to inactive.",
			//       "properties": {
			//         "Action": {
			//           "enum": [
			//             "DEACTIVATE"
			//           ],
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "Action"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The set of parameters for this mitigation action. You can specify only one type of parameter (in other words, you can apply only one action for each defined mitigation action).",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"add_things_to_thing_group_params": {
						// Property: AddThingsToThingGroupParams
						Description: "Parameters to define a mitigation action that moves devices associated with a certificate to one or more specified thing groups, typically for quarantine.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"override_dynamic_groups": {
									// Property: OverrideDynamicGroups
									Description: "Specifies if this mitigation action can move the things that triggered the mitigation action out of one or more dynamic thing groups.",
									Type:        types.BoolType,
									Computed:    true,
								},
								"thing_group_names": {
									// Property: ThingGroupNames
									Description: "The list of groups to which you want to add the things that triggered the mitigation action.",
									Type:        types.SetType{ElemType: types.StringType},
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"enable_io_t_logging_params": {
						// Property: EnableIoTLoggingParams
						Description: "Parameters to define a mitigation action that enables AWS IoT logging at a specified level of detail.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"log_level": {
									// Property: LogLevel
									Description: " Specifies which types of information are logged.",
									Type:        types.StringType,
									Computed:    true,
								},
								"role_arn_for_logging": {
									// Property: RoleArnForLogging
									Description: " The ARN of the IAM role used for logging.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"publish_finding_to_sns_params": {
						// Property: PublishFindingToSnsParams
						Description: "Parameters, to define a mitigation action that publishes findings to Amazon SNS. You can implement your own custom actions in response to the Amazon SNS messages.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"topic_arn": {
									// Property: TopicArn
									Description: "The ARN of the topic to which you want to publish the findings.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"replace_default_policy_version_params": {
						// Property: ReplaceDefaultPolicyVersionParams
						Description: "Parameters to define a mitigation action that adds a blank policy to restrict permissions.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"template_name": {
									// Property: TemplateName
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"update_ca_certificate_params": {
						// Property: UpdateCACertificateParams
						Description: "Parameters to define a mitigation action that changes the state of the CA certificate to inactive.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"action": {
									// Property: Action
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"update_device_certificate_params": {
						// Property: UpdateDeviceCertificateParams
						Description: "Parameters to define a mitigation action that changes the state of the device certificate to inactive.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"action": {
									// Property: Action
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"mitigation_action_arn": {
			// Property: MitigationActionArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"mitigation_action_id": {
			// Property: MitigationActionId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The tag's key.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The tag's value.",
			//         "maxLength": 256,
			//         "minLength": 1,
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
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The tag's key.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The tag's value.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::IoT::MitigationAction",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::MitigationAction").WithTerraformTypeName("awscc_iot_mitigation_action")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                                "Action",
		"action_name":                           "ActionName",
		"action_params":                         "ActionParams",
		"add_things_to_thing_group_params":      "AddThingsToThingGroupParams",
		"enable_io_t_logging_params":            "EnableIoTLoggingParams",
		"key":                                   "Key",
		"log_level":                             "LogLevel",
		"mitigation_action_arn":                 "MitigationActionArn",
		"mitigation_action_id":                  "MitigationActionId",
		"override_dynamic_groups":               "OverrideDynamicGroups",
		"publish_finding_to_sns_params":         "PublishFindingToSnsParams",
		"replace_default_policy_version_params": "ReplaceDefaultPolicyVersionParams",
		"role_arn":                              "RoleArn",
		"role_arn_for_logging":                  "RoleArnForLogging",
		"tags":                                  "Tags",
		"template_name":                         "TemplateName",
		"thing_group_names":                     "ThingGroupNames",
		"topic_arn":                             "TopicArn",
		"update_ca_certificate_params":          "UpdateCACertificateParams",
		"update_device_certificate_params":      "UpdateDeviceCertificateParams",
		"value":                                 "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
