// Code generated by generators/resource/main.go; DO NOT EDIT.

package backup

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_backup_report_plan", reportPlanResourceType)
}

// reportPlanResourceType returns the Terraform awscc_backup_report_plan resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Backup::ReportPlan resource type.
func reportPlanResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"report_delivery_channel": {
			// Property: ReportDeliveryChannel
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A structure that contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports.",
			//   "properties": {
			//     "Formats": {
			//       "description": "A list of the format of your reports: CSV, JSON, or both. If not specified, the default format is CSV.",
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "S3BucketName": {
			//       "description": "The unique name of the S3 bucket that receives your reports.",
			//       "type": "string"
			//     },
			//     "S3KeyPrefix": {
			//       "description": "The prefix for where AWS Backup Audit Manager delivers your reports to Amazon S3. The prefix is this part of the following path: s3://your-bucket-name/prefix/Backup/us-west-2/year/month/day/report-name. If not specified, there is no prefix.",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "S3BucketName"
			//   ],
			//   "type": "object"
			// }
			Description: "A structure that contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"formats": {
						// Property: Formats
						Description: "A list of the format of your reports: CSV, JSON, or both. If not specified, the default format is CSV.",
						Type:        types.SetType{ElemType: types.StringType},
						Optional:    true,
					},
					"s3_bucket_name": {
						// Property: S3BucketName
						Description: "The unique name of the S3 bucket that receives your reports.",
						Type:        types.StringType,
						Required:    true,
					},
					"s3_key_prefix": {
						// Property: S3KeyPrefix
						Description: "The prefix for where AWS Backup Audit Manager delivers your reports to Amazon S3. The prefix is this part of the following path: s3://your-bucket-name/prefix/Backup/us-west-2/year/month/day/report-name. If not specified, there is no prefix.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Required: true,
		},
		"report_plan_arn": {
			// Property: ReportPlanArn
			// CloudFormation resource type schema:
			// {
			//   "description": "An Amazon Resource Name (ARN) that uniquely identifies a resource. The format of the ARN depends on the resource type.",
			//   "type": "string"
			// }
			Description: "An Amazon Resource Name (ARN) that uniquely identifies a resource. The format of the ARN depends on the resource type.",
			Type:        types.StringType,
			Computed:    true,
		},
		"report_plan_description": {
			// Property: ReportPlanDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "An optional description of the report plan with a maximum of 1,024 characters.",
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "An optional description of the report plan with a maximum of 1,024 characters.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 1024),
			},
		},
		"report_plan_name": {
			// Property: ReportPlanName
			// CloudFormation resource type schema:
			// {
			//   "description": "The unique name of the report plan. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The unique name of the report plan. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"report_plan_tags": {
			// Property: ReportPlanTags
			// CloudFormation resource type schema:
			// {
			//   "description": "Metadata that you can assign to help organize the report plans that you create. Each tag is a key-value pair.",
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
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "Metadata that you can assign to help organize the report plans that you create. Each tag is a key-value pair.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Optional:    true,
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
		"report_setting": {
			// Property: ReportSetting
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Identifies the report template for the report. Reports are built using a report template.",
			//   "properties": {
			//     "FrameworkArns": {
			//       "description": "The Amazon Resource Names (ARNs) of the frameworks a report covers.",
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "ReportTemplate": {
			//       "description": "Identifies the report template for the report. Reports are built using a report template. The report templates are: `BACKUP_JOB_REPORT | COPY_JOB_REPORT | RESTORE_JOB_REPORT`",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "ReportTemplate"
			//   ],
			//   "type": "object"
			// }
			Description: "Identifies the report template for the report. Reports are built using a report template.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"framework_arns": {
						// Property: FrameworkArns
						Description: "The Amazon Resource Names (ARNs) of the frameworks a report covers.",
						Type:        types.SetType{ElemType: types.StringType},
						Optional:    true,
					},
					"report_template": {
						// Property: ReportTemplate
						Description: "Identifies the report template for the report. Reports are built using a report template. The report templates are: `BACKUP_JOB_REPORT | COPY_JOB_REPORT | RESTORE_JOB_REPORT`",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Required: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Contains detailed information about a report plan in AWS Backup Audit Manager.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Backup::ReportPlan").WithTerraformTypeName("awscc_backup_report_plan")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"formats":                 "Formats",
		"framework_arns":          "FrameworkArns",
		"key":                     "Key",
		"report_delivery_channel": "ReportDeliveryChannel",
		"report_plan_arn":         "ReportPlanArn",
		"report_plan_description": "ReportPlanDescription",
		"report_plan_name":        "ReportPlanName",
		"report_plan_tags":        "ReportPlanTags",
		"report_setting":          "ReportSetting",
		"report_template":         "ReportTemplate",
		"s3_bucket_name":          "S3BucketName",
		"s3_key_prefix":           "S3KeyPrefix",
		"value":                   "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
