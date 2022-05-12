// Code generated by generators/resource/main.go; DO NOT EDIT.

package lambda

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
	registry.AddResourceTypeFactory("awscc_lambda_code_signing_config", codeSigningConfigResourceType)
}

// codeSigningConfigResourceType returns the Terraform awscc_lambda_code_signing_config resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Lambda::CodeSigningConfig resource type.
func codeSigningConfigResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"allowed_publishers": {
			// Property: AllowedPublishers
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list",
			//   "properties": {
			//     "SigningProfileVersionArns": {
			//       "description": "List of Signing profile version Arns",
			//       "items": {
			//         "maxLength": 1024,
			//         "minLength": 12,
			//         "pattern": "arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\\-])+:([a-z]{2}(-gov)?-[a-z]+-\\d{1})?:(\\d{12})?:(.*)",
			//         "type": "string"
			//       },
			//       "maxItems": 20,
			//       "minItems": 1,
			//       "type": "array"
			//     }
			//   },
			//   "required": [
			//     "SigningProfileVersionArns"
			//   ],
			//   "type": "object"
			// }
			Description: "When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"signing_profile_version_arns": {
						// Property: SigningProfileVersionArns
						Description: "List of Signing profile version Arns",
						Type:        types.ListType{ElemType: types.StringType},
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(1, 20),
							validate.ArrayForEach(validate.StringLenBetween(12, 1024)),
							validate.ArrayForEach(validate.StringMatch(regexp.MustCompile("arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\\-])+:([a-z]{2}(-gov)?-[a-z]+-\\d{1})?:(\\d{12})?:(.*)"), "")),
						},
					},
				},
			),
			Required: true,
		},
		"code_signing_config_arn": {
			// Property: CodeSigningConfigArn
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique Arn for CodeSigningConfig resource",
			//   "pattern": "arn:(aws[a-zA-Z-]*)?:lambda:[a-z]{2}((-gov)|(-iso(b?)))?-[a-z]+-\\d{1}:\\d{12}:code-signing-config:csc-[a-z0-9]{17}",
			//   "type": "string"
			// }
			Description: "A unique Arn for CodeSigningConfig resource",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"code_signing_config_id": {
			// Property: CodeSigningConfigId
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique identifier for CodeSigningConfig resource",
			//   "pattern": "csc-[a-zA-Z0-9-_\\.]{17}",
			//   "type": "string"
			// }
			Description: "A unique identifier for CodeSigningConfig resource",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"code_signing_policies": {
			// Property: CodeSigningPolicies
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Policies to control how to act if a signature is invalid",
			//   "properties": {
			//     "UntrustedArtifactOnDeployment": {
			//       "default": "Warn",
			//       "description": "Indicates how Lambda operations involve updating the code artifact will operate. Default to Warn if not provided",
			//       "enum": [
			//         "Warn",
			//         "Enforce"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "UntrustedArtifactOnDeployment"
			//   ],
			//   "type": "object"
			// }
			Description: "Policies to control how to act if a signature is invalid",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"untrusted_artifact_on_deployment": {
						// Property: UntrustedArtifactOnDeployment
						Description: "Indicates how Lambda operations involve updating the code artifact will operate. Default to Warn if not provided",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"Warn",
								"Enforce",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							DefaultValue(types.String{Value: "Warn"}),
							tfsdk.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the CodeSigningConfig",
			//   "maxLength": 256,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "A description of the CodeSigningConfig",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 256),
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
		Description: "Resource Type definition for AWS::Lambda::CodeSigningConfig.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lambda::CodeSigningConfig").WithTerraformTypeName("awscc_lambda_code_signing_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allowed_publishers":               "AllowedPublishers",
		"code_signing_config_arn":          "CodeSigningConfigArn",
		"code_signing_config_id":           "CodeSigningConfigId",
		"code_signing_policies":            "CodeSigningPolicies",
		"description":                      "Description",
		"signing_profile_version_arns":     "SigningProfileVersionArns",
		"untrusted_artifact_on_deployment": "UntrustedArtifactOnDeployment",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
