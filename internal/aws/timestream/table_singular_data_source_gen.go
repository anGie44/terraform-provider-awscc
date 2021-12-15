// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package timestream

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_timestream_table", tableDataSourceType)
}

// tableDataSourceType returns the Terraform awscc_timestream_table data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Timestream::Table resource type.
func tableDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"database_name": {
			// Property: DatabaseName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the database which the table to be created belongs to.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name for the database which the table to be created belongs to.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The table name exposed as a read-only attribute.",
			//   "type": "string"
			// }
			Description: "The table name exposed as a read-only attribute.",
			Type:        types.StringType,
			Computed:    true,
		},
		"retention_properties": {
			// Property: RetentionProperties
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The retention duration of the memory store and the magnetic store.",
			//   "properties": {
			//     "MagneticStoreRetentionPeriodInDays": {
			//       "description": "The duration for which data must be stored in the magnetic store.",
			//       "type": "string"
			//     },
			//     "MemoryStoreRetentionPeriodInHours": {
			//       "description": "The duration for which data must be stored in the memory store.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The retention duration of the memory store and the magnetic store.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"magnetic_store_retention_period_in_days": {
						// Property: MagneticStoreRetentionPeriodInDays
						Description: "The duration for which data must be stored in the magnetic store.",
						Type:        types.StringType,
						Computed:    true,
					},
					"memory_store_retention_period_in_hours": {
						// Property: MemoryStoreRetentionPeriodInHours
						Description: "The duration for which data must be stored in the memory store.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"table_name": {
			// Property: TableName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name.",
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
			//     "description": "You can use the Resource Tags property to apply tags to resources, which can help you identify and categorize those resources.",
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
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "type": "array"
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
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
		Description: "Data Source schema for AWS::Timestream::Table",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Timestream::Table").WithTerraformTypeName("awscc_timestream_table")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":           "Arn",
		"database_name": "DatabaseName",
		"key":           "Key",
		"magnetic_store_retention_period_in_days": "MagneticStoreRetentionPeriodInDays",
		"memory_store_retention_period_in_hours":  "MemoryStoreRetentionPeriodInHours",
		"name":                                    "Name",
		"retention_properties":                    "RetentionProperties",
		"table_name":                              "TableName",
		"tags":                                    "Tags",
		"value":                                   "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
