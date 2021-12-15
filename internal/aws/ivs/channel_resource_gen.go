// Code generated by generators/resource/main.go; DO NOT EDIT.

package ivs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_ivs_channel", channelResourceType)
}

// channelResourceType returns the Terraform awscc_ivs_channel resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IVS::Channel resource type.
func channelResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Channel ARN is automatically generated on creation and assigned as the unique identifier.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Channel ARN is automatically generated on creation and assigned as the unique identifier.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"authorized": {
			// Property: Authorized
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether the channel is authorized.",
			//   "type": "boolean"
			// }
			Description: "Whether the channel is authorized.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"ingest_endpoint": {
			// Property: IngestEndpoint
			// CloudFormation resource type schema:
			// {
			//   "description": "Channel ingest endpoint, part of the definition of an ingest server, used when you set up streaming software.",
			//   "type": "string"
			// }
			Description: "Channel ingest endpoint, part of the definition of an ingest server, used when you set up streaming software.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"latency_mode": {
			// Property: LatencyMode
			// CloudFormation resource type schema:
			// {
			//   "description": "Channel latency mode.",
			//   "enum": [
			//     "NORMAL",
			//     "LOW"
			//   ],
			//   "type": "string"
			// }
			Description: "Channel latency mode.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"NORMAL",
					"LOW",
				}),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Channel",
			//   "maxLength": 128,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Channel",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 128),
			},
		},
		"playback_url": {
			// Property: PlaybackUrl
			// CloudFormation resource type schema:
			// {
			//   "description": "Channel Playback URL.",
			//   "type": "string"
			// }
			Description: "Channel Playback URL.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"recording_configuration_arn": {
			// Property: RecordingConfigurationArn
			// CloudFormation resource type schema:
			// {
			//   "default": "",
			//   "description": "Recording Configuration ARN. A value other than an empty string indicates that recording is enabled. Default: ?? (recording is disabled).",
			//   "maxLength": 128,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Recording Configuration ARN. A value other than an empty string indicates that recording is enabled. Default: ?? (recording is disabled).",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 128),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.String{Value: ""}),
				tfsdk.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the asset model.",
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
			Description: "A list of key-value pairs that contain metadata for the asset model.",
			Attributes: tfsdk.SetNestedAttributes(
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
							validate.StringLenBetween(1, 256),
						},
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
			},
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "Channel type, which determines the allowable resolution and bitrate. If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.",
			//   "enum": [
			//     "STANDARD",
			//     "BASIC"
			//   ],
			//   "type": "string"
			// }
			Description: "Channel type, which determines the allowable resolution and bitrate. If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"STANDARD",
					"BASIC",
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
		Description: "Resource Type definition for AWS::IVS::Channel",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IVS::Channel").WithTerraformTypeName("awscc_ivs_channel")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"authorized":                  "Authorized",
		"ingest_endpoint":             "IngestEndpoint",
		"key":                         "Key",
		"latency_mode":                "LatencyMode",
		"name":                        "Name",
		"playback_url":                "PlaybackUrl",
		"recording_configuration_arn": "RecordingConfigurationArn",
		"tags":                        "Tags",
		"type":                        "Type",
		"value":                       "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
