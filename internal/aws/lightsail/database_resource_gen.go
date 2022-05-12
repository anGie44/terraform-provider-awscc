// Code generated by generators/resource/main.go; DO NOT EDIT.

package lightsail

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
	registry.AddResourceTypeFactory("awscc_lightsail_database", databaseResourceType)
}

// databaseResourceType returns the Terraform awscc_lightsail_database resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Lightsail::Database resource type.
func databaseResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"availability_zone": {
			// Property: AvailabilityZone
			// CloudFormation resource type schema:
			// {
			//   "description": "The Availability Zone in which to create your new database. Use the us-east-2a case-sensitive format.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The Availability Zone in which to create your new database. Use the us-east-2a case-sensitive format.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"backup_retention": {
			// Property: BackupRetention
			// CloudFormation resource type schema:
			// {
			//   "description": "When true, enables automated backup retention for your database. Updates are applied during the next maintenance window because this can result in an outage.",
			//   "type": "boolean"
			// }
			Description: "When true, enables automated backup retention for your database. Updates are applied during the next maintenance window because this can result in an outage.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"ca_certificate_identifier": {
			// Property: CaCertificateIdentifier
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates the certificate that needs to be associated with the database.",
			//   "type": "string"
			// }
			Description: "Indicates the certificate that needs to be associated with the database.",
			Type:        types.StringType,
			Optional:    true,
		},
		"database_arn": {
			// Property: DatabaseArn
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
		"master_database_name": {
			// Property: MasterDatabaseName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the database to create when the Lightsail database resource is created. For MySQL, if this parameter isn't specified, no database is created in the database resource. For PostgreSQL, if this parameter isn't specified, a database named postgres is created in the database resource.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the database to create when the Lightsail database resource is created. For MySQL, if this parameter isn't specified, no database is created in the database resource. For PostgreSQL, if this parameter isn't specified, a database named postgres is created in the database resource.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"master_user_password": {
			// Property: MasterUserPassword
			// CloudFormation resource type schema:
			// {
			//   "description": "The password for the master user. The password can include any printable ASCII character except \"/\", \"\"\", or \"@\". It cannot contain spaces.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The password for the master user. The password can include any printable ASCII character except \"/\", \"\"\", or \"@\". It cannot contain spaces.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 63),
			},
			// MasterUserPassword is a write-only property.
		},
		"master_username": {
			// Property: MasterUsername
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the master user.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name for the master user.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 63),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"preferred_backup_window": {
			// Property: PreferredBackupWindow
			// CloudFormation resource type schema:
			// {
			//   "description": "The daily time range during which automated backups are created for your new database if automated backups are enabled.",
			//   "type": "string"
			// }
			Description: "The daily time range during which automated backups are created for your new database if automated backups are enabled.",
			Type:        types.StringType,
			Optional:    true,
		},
		"preferred_maintenance_window": {
			// Property: PreferredMaintenanceWindow
			// CloudFormation resource type schema:
			// {
			//   "description": "The weekly time range during which system maintenance can occur on your new database.",
			//   "type": "string"
			// }
			Description: "The weekly time range during which system maintenance can occur on your new database.",
			Type:        types.StringType,
			Optional:    true,
		},
		"publicly_accessible": {
			// Property: PubliclyAccessible
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the accessibility options for your new database. A value of true specifies a database that is available to resources outside of your Lightsail account. A value of false specifies a database that is available only to your Lightsail resources in the same region as your database.",
			//   "type": "boolean"
			// }
			Description: "Specifies the accessibility options for your new database. A value of true specifies a database that is available to resources outside of your Lightsail account. A value of false specifies a database that is available only to your Lightsail resources in the same region as your database.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"relational_database_blueprint_id": {
			// Property: RelationalDatabaseBlueprintId
			// CloudFormation resource type schema:
			// {
			//   "description": "The blueprint ID for your new database. A blueprint describes the major engine version of a database.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The blueprint ID for your new database. A blueprint describes the major engine version of a database.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"relational_database_bundle_id": {
			// Property: RelationalDatabaseBundleId
			// CloudFormation resource type schema:
			// {
			//   "description": "The bundle ID for your new database. A bundle describes the performance specifications for your database.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The bundle ID for your new database. A bundle describes the performance specifications for your database.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"relational_database_name": {
			// Property: RelationalDatabaseName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name to use for your new Lightsail database resource.",
			//   "maxLength": 255,
			//   "minLength": 2,
			//   "pattern": "\\w[\\w\\-]*\\w",
			//   "type": "string"
			// }
			Description: "The name to use for your new Lightsail database resource.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(2, 255),
				validate.StringMatch(regexp.MustCompile("\\w[\\w\\-]*\\w"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"relational_database_parameters": {
			// Property: RelationalDatabaseParameters
			// CloudFormation resource type schema:
			// {
			//   "description": "Update one or more parameters of the relational database.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Describes the parameters of the database.",
			//     "properties": {
			//       "AllowedValues": {
			//         "description": "Specifies the valid range of values for the parameter.",
			//         "type": "string"
			//       },
			//       "ApplyMethod": {
			//         "description": "Indicates when parameter updates are applied. Can be immediate or pending-reboot.",
			//         "type": "string"
			//       },
			//       "ApplyType": {
			//         "description": "Specifies the engine-specific parameter type.",
			//         "type": "string"
			//       },
			//       "DataType": {
			//         "description": "Specifies the valid data type for the parameter.",
			//         "type": "string"
			//       },
			//       "Description": {
			//         "description": "Provides a description of the parameter.",
			//         "type": "string"
			//       },
			//       "IsModifiable": {
			//         "description": "A Boolean value indicating whether the parameter can be modified.",
			//         "type": "boolean"
			//       },
			//       "ParameterName": {
			//         "description": "Specifies the name of the parameter.",
			//         "type": "string"
			//       },
			//       "ParameterValue": {
			//         "description": "Specifies the value of the parameter.",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Update one or more parameters of the relational database.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"allowed_values": {
						// Property: AllowedValues
						Description: "Specifies the valid range of values for the parameter.",
						Type:        types.StringType,
						Optional:    true,
					},
					"apply_method": {
						// Property: ApplyMethod
						Description: "Indicates when parameter updates are applied. Can be immediate or pending-reboot.",
						Type:        types.StringType,
						Optional:    true,
					},
					"apply_type": {
						// Property: ApplyType
						Description: "Specifies the engine-specific parameter type.",
						Type:        types.StringType,
						Optional:    true,
					},
					"data_type": {
						// Property: DataType
						Description: "Specifies the valid data type for the parameter.",
						Type:        types.StringType,
						Optional:    true,
					},
					"description": {
						// Property: Description
						Description: "Provides a description of the parameter.",
						Type:        types.StringType,
						Optional:    true,
					},
					"is_modifiable": {
						// Property: IsModifiable
						Description: "A Boolean value indicating whether the parameter can be modified.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"parameter_name": {
						// Property: ParameterName
						Description: "Specifies the name of the parameter.",
						Type:        types.StringType,
						Optional:    true,
					},
					"parameter_value": {
						// Property: ParameterValue
						Description: "Specifies the value of the parameter.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Optional: true,
			// RelationalDatabaseParameters is a write-only property.
		},
		"rotate_master_user_password": {
			// Property: RotateMasterUserPassword
			// CloudFormation resource type schema:
			// {
			//   "description": "When true, the master user password is changed to a new strong password generated by Lightsail. Use the get relational database master user password operation to get the new password.",
			//   "type": "boolean"
			// }
			Description: "When true, the master user password is changed to a new strong password generated by Lightsail. Use the get relational database master user password operation to get the new password.",
			Type:        types.BoolType,
			Optional:    true,
			// RotateMasterUserPassword is a write-only property.
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
			//       "Key"
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
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Optional: true,
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
		Description: "Resource Type definition for AWS::Lightsail::Database",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lightsail::Database").WithTerraformTypeName("awscc_lightsail_database")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allowed_values":                   "AllowedValues",
		"apply_method":                     "ApplyMethod",
		"apply_type":                       "ApplyType",
		"availability_zone":                "AvailabilityZone",
		"backup_retention":                 "BackupRetention",
		"ca_certificate_identifier":        "CaCertificateIdentifier",
		"data_type":                        "DataType",
		"database_arn":                     "DatabaseArn",
		"description":                      "Description",
		"is_modifiable":                    "IsModifiable",
		"key":                              "Key",
		"master_database_name":             "MasterDatabaseName",
		"master_user_password":             "MasterUserPassword",
		"master_username":                  "MasterUsername",
		"parameter_name":                   "ParameterName",
		"parameter_value":                  "ParameterValue",
		"preferred_backup_window":          "PreferredBackupWindow",
		"preferred_maintenance_window":     "PreferredMaintenanceWindow",
		"publicly_accessible":              "PubliclyAccessible",
		"relational_database_blueprint_id": "RelationalDatabaseBlueprintId",
		"relational_database_bundle_id":    "RelationalDatabaseBundleId",
		"relational_database_name":         "RelationalDatabaseName",
		"relational_database_parameters":   "RelationalDatabaseParameters",
		"rotate_master_user_password":      "RotateMasterUserPassword",
		"tags":                             "Tags",
		"value":                            "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/MasterUserPassword",
		"/properties/RelationalDatabaseParameters",
		"/properties/RotateMasterUserPassword",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
