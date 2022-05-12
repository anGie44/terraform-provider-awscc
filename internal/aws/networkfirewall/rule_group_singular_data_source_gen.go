// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package networkfirewall

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_networkfirewall_rule_group", ruleGroupDataSourceType)
}

// ruleGroupDataSourceType returns the Terraform awscc_networkfirewall_rule_group data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::NetworkFirewall::RuleGroup resource type.
func ruleGroupDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"capacity": {
			// Property: Capacity
			// CloudFormation resource type schema:
			// {
			//   "type": "integer"
			// }
			Type:     types.Int64Type,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 512,
			//   "minLength": 1,
			//   "pattern": "^.*$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"rule_group": {
			// Property: RuleGroup
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "RuleVariables": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "IPSets": {
			//           "additionalProperties": false,
			//           "patternProperties": {
			//             "": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Definition": {
			//                   "insertionOrder": true,
			//                   "items": {
			//                     "minLength": 1,
			//                     "pattern": "^.*$",
			//                     "type": "string"
			//                   },
			//                   "type": "array",
			//                   "uniqueItems": false
			//                 }
			//               },
			//               "type": "object"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "PortSets": {
			//           "additionalProperties": false,
			//           "patternProperties": {
			//             "": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Definition": {
			//                   "insertionOrder": true,
			//                   "items": {
			//                     "minLength": 1,
			//                     "pattern": "^.*$",
			//                     "type": "string"
			//                   },
			//                   "type": "array",
			//                   "uniqueItems": false
			//                 }
			//               },
			//               "type": "object"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "RulesSource": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "RulesSourceList": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "GeneratedRulesType": {
			//               "enum": [
			//                 "ALLOWLIST",
			//                 "DENYLIST"
			//               ],
			//               "type": "string"
			//             },
			//             "TargetTypes": {
			//               "insertionOrder": true,
			//               "items": {
			//                 "enum": [
			//                   "TLS_SNI",
			//                   "HTTP_HOST"
			//                 ],
			//                 "type": "string"
			//               },
			//               "type": "array",
			//               "uniqueItems": false
			//             },
			//             "Targets": {
			//               "insertionOrder": true,
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array",
			//               "uniqueItems": false
			//             }
			//           },
			//           "required": [
			//             "Targets",
			//             "TargetTypes",
			//             "GeneratedRulesType"
			//           ],
			//           "type": "object"
			//         },
			//         "RulesString": {
			//           "maxLength": 1000000,
			//           "minLength": 0,
			//           "type": "string"
			//         },
			//         "StatefulRules": {
			//           "insertionOrder": true,
			//           "items": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "Action": {
			//                 "enum": [
			//                   "PASS",
			//                   "DROP",
			//                   "ALERT"
			//                 ],
			//                 "type": "string"
			//               },
			//               "Header": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "Destination": {
			//                     "maxLength": 1024,
			//                     "minLength": 1,
			//                     "pattern": "^.*$",
			//                     "type": "string"
			//                   },
			//                   "DestinationPort": {
			//                     "maxLength": 1024,
			//                     "minLength": 1,
			//                     "pattern": "^.*$",
			//                     "type": "string"
			//                   },
			//                   "Direction": {
			//                     "enum": [
			//                       "FORWARD",
			//                       "ANY"
			//                     ],
			//                     "type": "string"
			//                   },
			//                   "Protocol": {
			//                     "enum": [
			//                       "IP",
			//                       "TCP",
			//                       "UDP",
			//                       "ICMP",
			//                       "HTTP",
			//                       "FTP",
			//                       "TLS",
			//                       "SMB",
			//                       "DNS",
			//                       "DCERPC",
			//                       "SSH",
			//                       "SMTP",
			//                       "IMAP",
			//                       "MSN",
			//                       "KRB5",
			//                       "IKEV2",
			//                       "TFTP",
			//                       "NTP",
			//                       "DHCP"
			//                     ],
			//                     "type": "string"
			//                   },
			//                   "Source": {
			//                     "maxLength": 1024,
			//                     "minLength": 1,
			//                     "pattern": "^.*$",
			//                     "type": "string"
			//                   },
			//                   "SourcePort": {
			//                     "maxLength": 1024,
			//                     "minLength": 1,
			//                     "pattern": "^.*$",
			//                     "type": "string"
			//                   }
			//                 },
			//                 "required": [
			//                   "Protocol",
			//                   "Source",
			//                   "SourcePort",
			//                   "Direction",
			//                   "Destination",
			//                   "DestinationPort"
			//                 ],
			//                 "type": "object"
			//               },
			//               "RuleOptions": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "additionalProperties": false,
			//                   "properties": {
			//                     "Keyword": {
			//                       "maxLength": 128,
			//                       "minLength": 1,
			//                       "pattern": "^.*$",
			//                       "type": "string"
			//                     },
			//                     "Settings": {
			//                       "insertionOrder": true,
			//                       "items": {
			//                         "maxLength": 8192,
			//                         "minLength": 1,
			//                         "pattern": "^.*$",
			//                         "type": "string"
			//                       },
			//                       "type": "array",
			//                       "uniqueItems": false
			//                     }
			//                   },
			//                   "required": [
			//                     "Keyword"
			//                   ],
			//                   "type": "object"
			//                 },
			//                 "type": "array",
			//                 "uniqueItems": false
			//               }
			//             },
			//             "required": [
			//               "Action",
			//               "Header",
			//               "RuleOptions"
			//             ],
			//             "type": "object"
			//           },
			//           "type": "array",
			//           "uniqueItems": false
			//         },
			//         "StatelessRulesAndCustomActions": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "CustomActions": {
			//               "insertionOrder": true,
			//               "items": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "ActionDefinition": {
			//                     "additionalProperties": false,
			//                     "properties": {
			//                       "PublishMetricAction": {
			//                         "additionalProperties": false,
			//                         "properties": {
			//                           "Dimensions": {
			//                             "insertionOrder": true,
			//                             "items": {
			//                               "additionalProperties": false,
			//                               "properties": {
			//                                 "Value": {
			//                                   "maxLength": 128,
			//                                   "minLength": 1,
			//                                   "pattern": "^[a-zA-Z0-9-_ ]+$",
			//                                   "type": "string"
			//                                 }
			//                               },
			//                               "required": [
			//                                 "Value"
			//                               ],
			//                               "type": "object"
			//                             },
			//                             "type": "array",
			//                             "uniqueItems": false
			//                           }
			//                         },
			//                         "required": [
			//                           "Dimensions"
			//                         ],
			//                         "type": "object"
			//                       }
			//                     },
			//                     "type": "object"
			//                   },
			//                   "ActionName": {
			//                     "maxLength": 128,
			//                     "minLength": 1,
			//                     "pattern": "^[a-zA-Z0-9]+$",
			//                     "type": "string"
			//                   }
			//                 },
			//                 "required": [
			//                   "ActionName",
			//                   "ActionDefinition"
			//                 ],
			//                 "type": "object"
			//               },
			//               "type": "array",
			//               "uniqueItems": false
			//             },
			//             "StatelessRules": {
			//               "insertionOrder": true,
			//               "items": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "Priority": {
			//                     "maximum": 65535,
			//                     "minimum": 1,
			//                     "type": "integer"
			//                   },
			//                   "RuleDefinition": {
			//                     "additionalProperties": false,
			//                     "properties": {
			//                       "Actions": {
			//                         "insertionOrder": true,
			//                         "items": {
			//                           "type": "string"
			//                         },
			//                         "type": "array",
			//                         "uniqueItems": false
			//                       },
			//                       "MatchAttributes": {
			//                         "additionalProperties": false,
			//                         "properties": {
			//                           "DestinationPorts": {
			//                             "insertionOrder": true,
			//                             "items": {
			//                               "additionalProperties": false,
			//                               "properties": {
			//                                 "FromPort": {
			//                                   "maximum": 65535,
			//                                   "minimum": 0,
			//                                   "type": "integer"
			//                                 },
			//                                 "ToPort": {
			//                                   "maximum": 65535,
			//                                   "minimum": 0,
			//                                   "type": "integer"
			//                                 }
			//                               },
			//                               "required": [
			//                                 "FromPort",
			//                                 "ToPort"
			//                               ],
			//                               "type": "object"
			//                             },
			//                             "type": "array",
			//                             "uniqueItems": false
			//                           },
			//                           "Destinations": {
			//                             "insertionOrder": true,
			//                             "items": {
			//                               "additionalProperties": false,
			//                               "properties": {
			//                                 "AddressDefinition": {
			//                                   "maxLength": 255,
			//                                   "minLength": 1,
			//                                   "pattern": "^([a-fA-F\\d:\\.]+/\\d{1,3})$",
			//                                   "type": "string"
			//                                 }
			//                               },
			//                               "required": [
			//                                 "AddressDefinition"
			//                               ],
			//                               "type": "object"
			//                             },
			//                             "type": "array",
			//                             "uniqueItems": false
			//                           },
			//                           "Protocols": {
			//                             "insertionOrder": true,
			//                             "items": {
			//                               "maximum": 255,
			//                               "minimum": 0,
			//                               "type": "integer"
			//                             },
			//                             "type": "array",
			//                             "uniqueItems": false
			//                           },
			//                           "SourcePorts": {
			//                             "insertionOrder": true,
			//                             "items": {
			//                               "additionalProperties": false,
			//                               "properties": {
			//                                 "FromPort": {
			//                                   "maximum": 65535,
			//                                   "minimum": 0,
			//                                   "type": "integer"
			//                                 },
			//                                 "ToPort": {
			//                                   "maximum": 65535,
			//                                   "minimum": 0,
			//                                   "type": "integer"
			//                                 }
			//                               },
			//                               "required": [
			//                                 "FromPort",
			//                                 "ToPort"
			//                               ],
			//                               "type": "object"
			//                             },
			//                             "type": "array",
			//                             "uniqueItems": false
			//                           },
			//                           "Sources": {
			//                             "insertionOrder": true,
			//                             "items": {
			//                               "additionalProperties": false,
			//                               "properties": {
			//                                 "AddressDefinition": {
			//                                   "maxLength": 255,
			//                                   "minLength": 1,
			//                                   "pattern": "^([a-fA-F\\d:\\.]+/\\d{1,3})$",
			//                                   "type": "string"
			//                                 }
			//                               },
			//                               "required": [
			//                                 "AddressDefinition"
			//                               ],
			//                               "type": "object"
			//                             },
			//                             "type": "array",
			//                             "uniqueItems": false
			//                           },
			//                           "TCPFlags": {
			//                             "insertionOrder": true,
			//                             "items": {
			//                               "additionalProperties": false,
			//                               "properties": {
			//                                 "Flags": {
			//                                   "insertionOrder": true,
			//                                   "items": {
			//                                     "enum": [
			//                                       "FIN",
			//                                       "SYN",
			//                                       "RST",
			//                                       "PSH",
			//                                       "ACK",
			//                                       "URG",
			//                                       "ECE",
			//                                       "CWR"
			//                                     ],
			//                                     "type": "string"
			//                                   },
			//                                   "type": "array",
			//                                   "uniqueItems": false
			//                                 },
			//                                 "Masks": {
			//                                   "insertionOrder": true,
			//                                   "items": {
			//                                     "enum": [
			//                                       "FIN",
			//                                       "SYN",
			//                                       "RST",
			//                                       "PSH",
			//                                       "ACK",
			//                                       "URG",
			//                                       "ECE",
			//                                       "CWR"
			//                                     ],
			//                                     "type": "string"
			//                                   },
			//                                   "type": "array",
			//                                   "uniqueItems": false
			//                                 }
			//                               },
			//                               "required": [
			//                                 "Flags"
			//                               ],
			//                               "type": "object"
			//                             },
			//                             "type": "array",
			//                             "uniqueItems": false
			//                           }
			//                         },
			//                         "type": "object"
			//                       }
			//                     },
			//                     "required": [
			//                       "MatchAttributes",
			//                       "Actions"
			//                     ],
			//                     "type": "object"
			//                   }
			//                 },
			//                 "required": [
			//                   "RuleDefinition",
			//                   "Priority"
			//                 ],
			//                 "type": "object"
			//               },
			//               "type": "array",
			//               "uniqueItems": false
			//             }
			//           },
			//           "required": [
			//             "StatelessRules"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "StatefulRuleOptions": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "RuleOrder": {
			//           "enum": [
			//             "DEFAULT_ACTION_ORDER",
			//             "STRICT_ORDER"
			//           ],
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "RulesSource"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"rule_variables": {
						// Property: RuleVariables
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"ip_sets": {
									// Property: IPSets
									// Pattern: ""
									Attributes: tfsdk.MapNestedAttributes(
										map[string]tfsdk.Attribute{
											"definition": {
												// Property: Definition
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
										},
										tfsdk.MapNestedAttributesOptions{},
									),
									Computed: true,
								},
								"port_sets": {
									// Property: PortSets
									// Pattern: ""
									Attributes: tfsdk.MapNestedAttributes(
										map[string]tfsdk.Attribute{
											"definition": {
												// Property: Definition
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
										},
										tfsdk.MapNestedAttributesOptions{},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"rules_source": {
						// Property: RulesSource
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"rules_source_list": {
									// Property: RulesSourceList
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"generated_rules_type": {
												// Property: GeneratedRulesType
												Type:     types.StringType,
												Computed: true,
											},
											"target_types": {
												// Property: TargetTypes
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"targets": {
												// Property: Targets
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"rules_string": {
									// Property: RulesString
									Type:     types.StringType,
									Computed: true,
								},
								"stateful_rules": {
									// Property: StatefulRules
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"action": {
												// Property: Action
												Type:     types.StringType,
												Computed: true,
											},
											"header": {
												// Property: Header
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"destination": {
															// Property: Destination
															Type:     types.StringType,
															Computed: true,
														},
														"destination_port": {
															// Property: DestinationPort
															Type:     types.StringType,
															Computed: true,
														},
														"direction": {
															// Property: Direction
															Type:     types.StringType,
															Computed: true,
														},
														"protocol": {
															// Property: Protocol
															Type:     types.StringType,
															Computed: true,
														},
														"source": {
															// Property: Source
															Type:     types.StringType,
															Computed: true,
														},
														"source_port": {
															// Property: SourcePort
															Type:     types.StringType,
															Computed: true,
														},
													},
												),
												Computed: true,
											},
											"rule_options": {
												// Property: RuleOptions
												Attributes: tfsdk.ListNestedAttributes(
													map[string]tfsdk.Attribute{
														"keyword": {
															// Property: Keyword
															Type:     types.StringType,
															Computed: true,
														},
														"settings": {
															// Property: Settings
															Type:     types.ListType{ElemType: types.StringType},
															Computed: true,
														},
													},
													tfsdk.ListNestedAttributesOptions{},
												),
												Computed: true,
											},
										},
										tfsdk.ListNestedAttributesOptions{},
									),
									Computed: true,
								},
								"stateless_rules_and_custom_actions": {
									// Property: StatelessRulesAndCustomActions
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"custom_actions": {
												// Property: CustomActions
												Attributes: tfsdk.ListNestedAttributes(
													map[string]tfsdk.Attribute{
														"action_definition": {
															// Property: ActionDefinition
															Attributes: tfsdk.SingleNestedAttributes(
																map[string]tfsdk.Attribute{
																	"publish_metric_action": {
																		// Property: PublishMetricAction
																		Attributes: tfsdk.SingleNestedAttributes(
																			map[string]tfsdk.Attribute{
																				"dimensions": {
																					// Property: Dimensions
																					Attributes: tfsdk.ListNestedAttributes(
																						map[string]tfsdk.Attribute{
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
																			},
																		),
																		Computed: true,
																	},
																},
															),
															Computed: true,
														},
														"action_name": {
															// Property: ActionName
															Type:     types.StringType,
															Computed: true,
														},
													},
													tfsdk.ListNestedAttributesOptions{},
												),
												Computed: true,
											},
											"stateless_rules": {
												// Property: StatelessRules
												Attributes: tfsdk.ListNestedAttributes(
													map[string]tfsdk.Attribute{
														"priority": {
															// Property: Priority
															Type:     types.Int64Type,
															Computed: true,
														},
														"rule_definition": {
															// Property: RuleDefinition
															Attributes: tfsdk.SingleNestedAttributes(
																map[string]tfsdk.Attribute{
																	"actions": {
																		// Property: Actions
																		Type:     types.ListType{ElemType: types.StringType},
																		Computed: true,
																	},
																	"match_attributes": {
																		// Property: MatchAttributes
																		Attributes: tfsdk.SingleNestedAttributes(
																			map[string]tfsdk.Attribute{
																				"destination_ports": {
																					// Property: DestinationPorts
																					Attributes: tfsdk.ListNestedAttributes(
																						map[string]tfsdk.Attribute{
																							"from_port": {
																								// Property: FromPort
																								Type:     types.Int64Type,
																								Computed: true,
																							},
																							"to_port": {
																								// Property: ToPort
																								Type:     types.Int64Type,
																								Computed: true,
																							},
																						},
																						tfsdk.ListNestedAttributesOptions{},
																					),
																					Computed: true,
																				},
																				"destinations": {
																					// Property: Destinations
																					Attributes: tfsdk.ListNestedAttributes(
																						map[string]tfsdk.Attribute{
																							"address_definition": {
																								// Property: AddressDefinition
																								Type:     types.StringType,
																								Computed: true,
																							},
																						},
																						tfsdk.ListNestedAttributesOptions{},
																					),
																					Computed: true,
																				},
																				"protocols": {
																					// Property: Protocols
																					Type:     types.ListType{ElemType: types.Int64Type},
																					Computed: true,
																				},
																				"source_ports": {
																					// Property: SourcePorts
																					Attributes: tfsdk.ListNestedAttributes(
																						map[string]tfsdk.Attribute{
																							"from_port": {
																								// Property: FromPort
																								Type:     types.Int64Type,
																								Computed: true,
																							},
																							"to_port": {
																								// Property: ToPort
																								Type:     types.Int64Type,
																								Computed: true,
																							},
																						},
																						tfsdk.ListNestedAttributesOptions{},
																					),
																					Computed: true,
																				},
																				"sources": {
																					// Property: Sources
																					Attributes: tfsdk.ListNestedAttributes(
																						map[string]tfsdk.Attribute{
																							"address_definition": {
																								// Property: AddressDefinition
																								Type:     types.StringType,
																								Computed: true,
																							},
																						},
																						tfsdk.ListNestedAttributesOptions{},
																					),
																					Computed: true,
																				},
																				"tcp_flags": {
																					// Property: TCPFlags
																					Attributes: tfsdk.ListNestedAttributes(
																						map[string]tfsdk.Attribute{
																							"flags": {
																								// Property: Flags
																								Type:     types.ListType{ElemType: types.StringType},
																								Computed: true,
																							},
																							"masks": {
																								// Property: Masks
																								Type:     types.ListType{ElemType: types.StringType},
																								Computed: true,
																							},
																						},
																						tfsdk.ListNestedAttributesOptions{},
																					),
																					Computed: true,
																				},
																			},
																		),
																		Computed: true,
																	},
																},
															),
															Computed: true,
														},
													},
													tfsdk.ListNestedAttributesOptions{},
												),
												Computed: true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"stateful_rule_options": {
						// Property: StatefulRuleOptions
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"rule_order": {
									// Property: RuleOrder
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"rule_group_arn": {
			// Property: RuleGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "A resource ARN.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^(arn:aws.*)$",
			//   "type": "string"
			// }
			Description: "A resource ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"rule_group_id": {
			// Property: RuleGroupId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 36,
			//   "minLength": 36,
			//   "pattern": "^([0-9a-f]{8})-([0-9a-f]{4}-){3}([0-9a-f]{12})$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"rule_group_name": {
			// Property: RuleGroupName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9-]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "^.*$",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 255,
			//         "minLength": 0,
			//         "pattern": "^.*$",
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
			Attributes: tfsdk.SetNestedAttributes(
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
				tfsdk.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "STATELESS",
			//     "STATEFUL"
			//   ],
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
		Description: "Data Source schema for AWS::NetworkFirewall::RuleGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkFirewall::RuleGroup").WithTerraformTypeName("awscc_networkfirewall_rule_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                             "Action",
		"action_definition":                  "ActionDefinition",
		"action_name":                        "ActionName",
		"actions":                            "Actions",
		"address_definition":                 "AddressDefinition",
		"capacity":                           "Capacity",
		"custom_actions":                     "CustomActions",
		"definition":                         "Definition",
		"description":                        "Description",
		"destination":                        "Destination",
		"destination_port":                   "DestinationPort",
		"destination_ports":                  "DestinationPorts",
		"destinations":                       "Destinations",
		"dimensions":                         "Dimensions",
		"direction":                          "Direction",
		"flags":                              "Flags",
		"from_port":                          "FromPort",
		"generated_rules_type":               "GeneratedRulesType",
		"header":                             "Header",
		"ip_sets":                            "IPSets",
		"key":                                "Key",
		"keyword":                            "Keyword",
		"masks":                              "Masks",
		"match_attributes":                   "MatchAttributes",
		"port_sets":                          "PortSets",
		"priority":                           "Priority",
		"protocol":                           "Protocol",
		"protocols":                          "Protocols",
		"publish_metric_action":              "PublishMetricAction",
		"rule_definition":                    "RuleDefinition",
		"rule_group":                         "RuleGroup",
		"rule_group_arn":                     "RuleGroupArn",
		"rule_group_id":                      "RuleGroupId",
		"rule_group_name":                    "RuleGroupName",
		"rule_options":                       "RuleOptions",
		"rule_order":                         "RuleOrder",
		"rule_variables":                     "RuleVariables",
		"rules_source":                       "RulesSource",
		"rules_source_list":                  "RulesSourceList",
		"rules_string":                       "RulesString",
		"settings":                           "Settings",
		"source":                             "Source",
		"source_port":                        "SourcePort",
		"source_ports":                       "SourcePorts",
		"sources":                            "Sources",
		"stateful_rule_options":              "StatefulRuleOptions",
		"stateful_rules":                     "StatefulRules",
		"stateless_rules":                    "StatelessRules",
		"stateless_rules_and_custom_actions": "StatelessRulesAndCustomActions",
		"tags":                               "Tags",
		"target_types":                       "TargetTypes",
		"targets":                            "Targets",
		"tcp_flags":                          "TCPFlags",
		"to_port":                            "ToPort",
		"type":                               "Type",
		"value":                              "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
