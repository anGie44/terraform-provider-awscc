// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_apigateway_domain_name", domainNameDataSourceType)
}

// domainNameDataSourceType returns the Terraform awscc_apigateway_domain_name data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ApiGateway::DomainName resource type.
func domainNameDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"certificate_arn": {
			// Property: CertificateArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"distribution_domain_name": {
			// Property: DistributionDomainName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"distribution_hosted_zone_id": {
			// Property: DistributionHostedZoneId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"endpoint_configuration": {
			// Property: EndpointConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Types": {
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"types": {
						// Property: Types
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"mutual_tls_authentication": {
			// Property: MutualTlsAuthentication
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "TruststoreUri": {
			//       "type": "string"
			//     },
			//     "TruststoreVersion": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"truststore_uri": {
						// Property: TruststoreUri
						Type:     types.StringType,
						Computed: true,
					},
					"truststore_version": {
						// Property: TruststoreVersion
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"ownership_verification_certificate_arn": {
			// Property: OwnershipVerificationCertificateArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"regional_certificate_arn": {
			// Property: RegionalCertificateArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"regional_domain_name": {
			// Property: RegionalDomainName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"regional_hosted_zone_id": {
			// Property: RegionalHostedZoneId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"security_policy": {
			// Property: SecurityPolicy
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
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
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
		Description: "Data Source schema for AWS::ApiGateway::DomainName",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::DomainName").WithTerraformTypeName("awscc_apigateway_domain_name")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"certificate_arn":                        "CertificateArn",
		"distribution_domain_name":               "DistributionDomainName",
		"distribution_hosted_zone_id":            "DistributionHostedZoneId",
		"domain_name":                            "DomainName",
		"endpoint_configuration":                 "EndpointConfiguration",
		"key":                                    "Key",
		"mutual_tls_authentication":              "MutualTlsAuthentication",
		"ownership_verification_certificate_arn": "OwnershipVerificationCertificateArn",
		"regional_certificate_arn":               "RegionalCertificateArn",
		"regional_domain_name":                   "RegionalDomainName",
		"regional_hosted_zone_id":                "RegionalHostedZoneId",
		"security_policy":                        "SecurityPolicy",
		"tags":                                   "Tags",
		"truststore_uri":                         "TruststoreUri",
		"truststore_version":                     "TruststoreVersion",
		"types":                                  "Types",
		"value":                                  "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
