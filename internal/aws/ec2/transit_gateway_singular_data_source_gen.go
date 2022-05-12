// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ec2_transit_gateway", transitGatewayDataSourceType)
}

// transitGatewayDataSourceType returns the Terraform awscc_ec2_transit_gateway data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::EC2::TransitGateway resource type.
func transitGatewayDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"amazon_side_asn": {
			// Property: AmazonSideAsn
			// CloudFormation resource type schema:
			// {
			//   "format": "int64",
			//   "type": "integer"
			// }
			Type:     types.Int64Type,
			Computed: true,
		},
		"association_default_route_table_id": {
			// Property: AssociationDefaultRouteTableId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"auto_accept_shared_attachments": {
			// Property: AutoAcceptSharedAttachments
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"default_route_table_association": {
			// Property: DefaultRouteTableAssociation
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"default_route_table_propagation": {
			// Property: DefaultRouteTablePropagation
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"dns_support": {
			// Property: DnsSupport
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"multicast_support": {
			// Property: MulticastSupport
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"propagation_default_route_table_id": {
			// Property: PropagationDefaultRouteTableId
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
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
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
		"transit_gateway_cidr_blocks": {
			// Property: TransitGatewayCidrBlocks
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"vpn_ecmp_support": {
			// Property: VpnEcmpSupport
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::EC2::TransitGateway",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::TransitGateway").WithTerraformTypeName("awscc_ec2_transit_gateway")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"amazon_side_asn":                    "AmazonSideAsn",
		"association_default_route_table_id": "AssociationDefaultRouteTableId",
		"auto_accept_shared_attachments":     "AutoAcceptSharedAttachments",
		"default_route_table_association":    "DefaultRouteTableAssociation",
		"default_route_table_propagation":    "DefaultRouteTablePropagation",
		"description":                        "Description",
		"dns_support":                        "DnsSupport",
		"id":                                 "Id",
		"key":                                "Key",
		"multicast_support":                  "MulticastSupport",
		"propagation_default_route_table_id": "PropagationDefaultRouteTableId",
		"tags":                               "Tags",
		"transit_gateway_cidr_blocks":        "TransitGatewayCidrBlocks",
		"value":                              "Value",
		"vpn_ecmp_support":                   "VpnEcmpSupport",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
