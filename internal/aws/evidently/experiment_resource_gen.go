// Code generated by generators/resource/main.go; DO NOT EDIT.

package evidently

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_evidently_experiment", experimentResourceType)
}

// experimentResourceType returns the Terraform awscc_evidently_experiment resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Evidently::Experiment resource type.
func experimentResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 160,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 160),
			},
		},
		"metric_goals": {
			// Property: MetricGoals
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "DesiredChange": {
			//         "enum": [
			//           "INCREASE",
			//           "DECREASE"
			//         ],
			//         "type": "string"
			//       },
			//       "EntityIdKey": {
			//         "description": "The JSON path to reference the entity id in the event.",
			//         "type": "string"
			//       },
			//       "EventPattern": {
			//         "description": "Event patterns have the same structure as the events they match. Rules use event patterns to select events. An event pattern either matches an event or it doesn't.",
			//         "type": "string"
			//       },
			//       "MetricName": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "UnitLabel": {
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "ValueKey": {
			//         "description": "The JSON path to reference the numerical metric value in the event.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "MetricName",
			//       "EntityIdKey",
			//       "ValueKey",
			//       "EventPattern",
			//       "DesiredChange"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 3,
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"desired_change": {
						// Property: DesiredChange
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"INCREASE",
								"DECREASE",
							}),
						},
					},
					"entity_id_key": {
						// Property: EntityIdKey
						Description: "The JSON path to reference the entity id in the event.",
						Type:        types.StringType,
						Required:    true,
					},
					"event_pattern": {
						// Property: EventPattern
						Description: "Event patterns have the same structure as the events they match. Rules use event patterns to select events. An event pattern either matches an event or it doesn't.",
						Type:        types.StringType,
						Required:    true,
					},
					"metric_name": {
						// Property: MetricName
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
					"unit_label": {
						// Property: UnitLabel
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
					},
					"value_key": {
						// Property: ValueKey
						Description: "The JSON path to reference the numerical metric value in the event.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(1, 3),
				validate.UniqueItems(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 127,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 127),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"online_ab_config": {
			// Property: OnlineAbConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "ControlTreatmentName": {
			//       "maxLength": 127,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "TreatmentWeights": {
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "SplitWeight": {
			//             "maximum": 100000,
			//             "minimum": 0,
			//             "type": "integer"
			//           },
			//           "Treatment": {
			//             "maxLength": 127,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Treatment",
			//           "SplitWeight"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"control_treatment_name": {
						// Property: ControlTreatmentName
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 127),
						},
					},
					"treatment_weights": {
						// Property: TreatmentWeights
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"split_weight": {
									// Property: SplitWeight
									Type:     types.NumberType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.IntBetween(0, 100000),
									},
								},
								"treatment": {
									// Property: Treatment
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 127),
									},
								},
							},
							tfsdk.SetNestedAttributesOptions{},
						),
						Optional: true,
					},
				},
			),
			Required: true,
		},
		"project": {
			// Property: Project
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 2048),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"randomization_salt": {
			// Property: RandomizationSalt
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 127,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 127),
			},
		},
		"sampling_rate": {
			// Property: SamplingRate
			// CloudFormation resource type schema:
			// {
			//   "maximum": 100000,
			//   "minimum": 0,
			//   "type": "integer"
			// }
			Type:     types.NumberType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.IntBetween(0, 100000),
			},
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
			//         "pattern": "",
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
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
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
				tfsdk.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
		"treatments": {
			// Property: Treatments
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Description": {
			//         "type": "string"
			//       },
			//       "Feature": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "TreatmentName": {
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Variation": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "TreatmentName",
			//       "Feature",
			//       "Variation"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 5,
			//   "minItems": 2,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"description": {
						// Property: Description
						Type:     types.StringType,
						Optional: true,
					},
					"feature": {
						// Property: Feature
						Type:     types.StringType,
						Required: true,
					},
					"treatment_name": {
						// Property: TreatmentName
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 127),
						},
					},
					"variation": {
						// Property: Variation
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(2, 5),
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
		Description: "Resource Type definition for AWS::Evidently::Experiment.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Evidently::Experiment").WithTerraformTypeName("awscc_evidently_experiment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"control_treatment_name": "ControlTreatmentName",
		"description":            "Description",
		"desired_change":         "DesiredChange",
		"entity_id_key":          "EntityIdKey",
		"event_pattern":          "EventPattern",
		"feature":                "Feature",
		"key":                    "Key",
		"metric_goals":           "MetricGoals",
		"metric_name":            "MetricName",
		"name":                   "Name",
		"online_ab_config":       "OnlineAbConfig",
		"project":                "Project",
		"randomization_salt":     "RandomizationSalt",
		"sampling_rate":          "SamplingRate",
		"split_weight":           "SplitWeight",
		"tags":                   "Tags",
		"treatment":              "Treatment",
		"treatment_name":         "TreatmentName",
		"treatment_weights":      "TreatmentWeights",
		"treatments":             "Treatments",
		"unit_label":             "UnitLabel",
		"value":                  "Value",
		"value_key":              "ValueKey",
		"variation":              "Variation",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
