// Code generated by generators/resource/main.go; DO NOT EDIT.

package frauddetector

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
	registry.AddResourceTypeFactory("awscc_frauddetector_variable", variableResourceType)
}

// variableResourceType returns the Terraform awscc_frauddetector_variable resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::FraudDetector::Variable resource type.
func variableResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the variable.",
			//   "type": "string"
			// }
			Description: "The ARN of the variable.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"created_time": {
			// Property: CreatedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The time when the variable was created.",
			//   "type": "string"
			// }
			Description: "The time when the variable was created.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"data_source": {
			// Property: DataSource
			// CloudFormation resource type schema:
			// {
			//   "description": "The source of the data.",
			//   "enum": [
			//     "EVENT",
			//     "EXTERNAL_MODEL_SCORE"
			//   ],
			//   "type": "string"
			// }
			Description: "The source of the data.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"EVENT",
					"EXTERNAL_MODEL_SCORE",
				}),
			},
		},
		"data_type": {
			// Property: DataType
			// CloudFormation resource type schema:
			// {
			//   "description": "The data type.",
			//   "enum": [
			//     "STRING",
			//     "INTEGER",
			//     "FLOAT",
			//     "BOOLEAN"
			//   ],
			//   "type": "string"
			// }
			Description: "The data type.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"STRING",
					"INTEGER",
					"FLOAT",
					"BOOLEAN",
				}),
			},
		},
		"default_value": {
			// Property: DefaultValue
			// CloudFormation resource type schema:
			// {
			//   "description": "The default value for the variable when no value is received.",
			//   "type": "string"
			// }
			Description: "The default value for the variable when no value is received.",
			Type:        types.StringType,
			Required:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The description.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
			},
		},
		"last_updated_time": {
			// Property: LastUpdatedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The time when the variable was last updated.",
			//   "type": "string"
			// }
			Description: "The time when the variable was last updated.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the variable.",
			//   "pattern": "^[a-z_][a-z0-9_]{0,99}?$",
			//   "type": "string"
			// }
			Description: "The name of the variable.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^[a-z_][a-z0-9_]{0,99}?$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Tags associated with this variable.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
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
			//   "maxItems": 200,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "Tags associated with this variable.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(200),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"variable_type": {
			// Property: VariableType
			// CloudFormation resource type schema:
			// {
			//   "description": "The variable type. For more information see https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types",
			//   "enum": [
			//     "AUTH_CODE",
			//     "AVS",
			//     "BILLING_ADDRESS_L1",
			//     "BILLING_ADDRESS_L2",
			//     "BILLING_CITY",
			//     "BILLING_COUNTRY",
			//     "BILLING_NAME",
			//     "BILLING_PHONE",
			//     "BILLING_STATE",
			//     "BILLING_ZIP",
			//     "CARD_BIN",
			//     "CATEGORICAL",
			//     "CURRENCY_CODE",
			//     "EMAIL_ADDRESS",
			//     "FINGERPRINT",
			//     "FRAUD_LABEL",
			//     "FREE_FORM_TEXT",
			//     "IP_ADDRESS",
			//     "NUMERIC",
			//     "ORDER_ID",
			//     "PAYMENT_TYPE",
			//     "PHONE_NUMBER",
			//     "PRICE",
			//     "PRODUCT_CATEGORY",
			//     "SHIPPING_ADDRESS_L1",
			//     "SHIPPING_ADDRESS_L2",
			//     "SHIPPING_CITY",
			//     "SHIPPING_COUNTRY",
			//     "SHIPPING_NAME",
			//     "SHIPPING_PHONE",
			//     "SHIPPING_STATE",
			//     "SHIPPING_ZIP",
			//     "USERAGENT"
			//   ],
			//   "type": "string"
			// }
			Description: "The variable type. For more information see https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"AUTH_CODE",
					"AVS",
					"BILLING_ADDRESS_L1",
					"BILLING_ADDRESS_L2",
					"BILLING_CITY",
					"BILLING_COUNTRY",
					"BILLING_NAME",
					"BILLING_PHONE",
					"BILLING_STATE",
					"BILLING_ZIP",
					"CARD_BIN",
					"CATEGORICAL",
					"CURRENCY_CODE",
					"EMAIL_ADDRESS",
					"FINGERPRINT",
					"FRAUD_LABEL",
					"FREE_FORM_TEXT",
					"IP_ADDRESS",
					"NUMERIC",
					"ORDER_ID",
					"PAYMENT_TYPE",
					"PHONE_NUMBER",
					"PRICE",
					"PRODUCT_CATEGORY",
					"SHIPPING_ADDRESS_L1",
					"SHIPPING_ADDRESS_L2",
					"SHIPPING_CITY",
					"SHIPPING_COUNTRY",
					"SHIPPING_NAME",
					"SHIPPING_PHONE",
					"SHIPPING_STATE",
					"SHIPPING_ZIP",
					"USERAGENT",
				}),
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
		Description: "A resource schema for a Variable in Amazon Fraud Detector.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::FraudDetector::Variable").WithTerraformTypeName("awscc_frauddetector_variable")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":               "Arn",
		"created_time":      "CreatedTime",
		"data_source":       "DataSource",
		"data_type":         "DataType",
		"default_value":     "DefaultValue",
		"description":       "Description",
		"key":               "Key",
		"last_updated_time": "LastUpdatedTime",
		"name":              "Name",
		"tags":              "Tags",
		"value":             "Value",
		"variable_type":     "VariableType",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
