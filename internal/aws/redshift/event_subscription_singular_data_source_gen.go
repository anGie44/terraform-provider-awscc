// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_redshift_event_subscription", eventSubscriptionDataSourceType)
}

// eventSubscriptionDataSourceType returns the Terraform awscc_redshift_event_subscription data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Redshift::EventSubscription resource type.
func eventSubscriptionDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"cust_subscription_id": {
			// Property: CustSubscriptionId
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Amazon Redshift event notification subscription.",
			//   "type": "string"
			// }
			Description: "The name of the Amazon Redshift event notification subscription.",
			Type:        types.StringType,
			Computed:    true,
		},
		"customer_aws_id": {
			// Property: CustomerAwsId
			// CloudFormation resource type schema:
			// {
			//   "description": "The AWS account associated with the Amazon Redshift event notification subscription.",
			//   "type": "string"
			// }
			Description: "The AWS account associated with the Amazon Redshift event notification subscription.",
			Type:        types.StringType,
			Computed:    true,
		},
		"enabled": {
			// Property: Enabled
			// CloudFormation resource type schema:
			// {
			//   "description": "A boolean value; set to true to activate the subscription, and set to false to create the subscription but not activate it.",
			//   "type": "boolean"
			// }
			Description: "A boolean value; set to true to activate the subscription, and set to false to create the subscription but not activate it.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"event_categories": {
			// Property: EventCategories
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the Amazon Redshift event categories to be published by the event notification subscription.",
			//   "insertionOrder": false,
			//   "items": {
			//     "enum": [
			//       "configuration",
			//       "management",
			//       "monitoring",
			//       "security",
			//       "pending"
			//     ],
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "Specifies the Amazon Redshift event categories to be published by the event notification subscription.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"event_categories_list": {
			// Property: EventCategoriesList
			// CloudFormation resource type schema:
			// {
			//   "description": "The list of Amazon Redshift event categories specified in the event notification subscription.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "The list of Amazon Redshift event categories specified in the event notification subscription.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"severity": {
			// Property: Severity
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the Amazon Redshift event severity to be published by the event notification subscription.",
			//   "enum": [
			//     "ERROR",
			//     "INFO"
			//   ],
			//   "type": "string"
			// }
			Description: "Specifies the Amazon Redshift event severity to be published by the event notification subscription.",
			Type:        types.StringType,
			Computed:    true,
		},
		"sns_topic_arn": {
			// Property: SnsTopicArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the Amazon SNS topic used to transmit the event notifications.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the Amazon SNS topic used to transmit the event notifications.",
			Type:        types.StringType,
			Computed:    true,
		},
		"source_ids": {
			// Property: SourceIds
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of one or more identifiers of Amazon Redshift source objects.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "A list of one or more identifiers of Amazon Redshift source objects.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"source_ids_list": {
			// Property: SourceIdsList
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of the sources that publish events to the Amazon Redshift event notification subscription.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "A list of the sources that publish events to the Amazon Redshift event notification subscription.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"source_type": {
			// Property: SourceType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of source that will be generating the events.",
			//   "enum": [
			//     "cluster",
			//     "cluster-parameter-group",
			//     "cluster-security-group",
			//     "cluster-snapshot",
			//     "scheduled-action"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of source that will be generating the events.",
			Type:        types.StringType,
			Computed:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of the Amazon Redshift event notification subscription.",
			//   "enum": [
			//     "active",
			//     "no-permission",
			//     "topic-not-exist"
			//   ],
			//   "type": "string"
			// }
			Description: "The status of the Amazon Redshift event notification subscription.",
			Type:        types.StringType,
			Computed:    true,
		},
		"subscription_creation_time": {
			// Property: SubscriptionCreationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The date and time the Amazon Redshift event notification subscription was created.",
			//   "type": "string"
			// }
			Description: "The date and time the Amazon Redshift event notification subscription was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"subscription_name": {
			// Property: SubscriptionName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Amazon Redshift event notification subscription",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the Amazon Redshift event notification subscription",
			Type:        types.StringType,
			Computed:    true,
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
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
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
		Description: "Data Source schema for AWS::Redshift::EventSubscription",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Redshift::EventSubscription").WithTerraformTypeName("awscc_redshift_event_subscription")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cust_subscription_id":       "CustSubscriptionId",
		"customer_aws_id":            "CustomerAwsId",
		"enabled":                    "Enabled",
		"event_categories":           "EventCategories",
		"event_categories_list":      "EventCategoriesList",
		"key":                        "Key",
		"severity":                   "Severity",
		"sns_topic_arn":              "SnsTopicArn",
		"source_ids":                 "SourceIds",
		"source_ids_list":            "SourceIdsList",
		"source_type":                "SourceType",
		"status":                     "Status",
		"subscription_creation_time": "SubscriptionCreationTime",
		"subscription_name":          "SubscriptionName",
		"tags":                       "Tags",
		"value":                      "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
