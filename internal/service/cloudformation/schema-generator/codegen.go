package generator

import (
	"fmt"
	"strings"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/iancoleman/strcase"
)

// Features of the generated code.
type Features int

const (
	UsesRegexp Features = 1 << iota
	UsesValidation
	HasUpdatableProperty
)

func RootPropertySchema(r *cfschema.Resource, name string) (string, Features) {
	if r == nil {
		return "", 0
	}

	property, ok := r.Properties[name]

	if !ok || property == nil {
		return "", 0
	}

	return PropertySchema(r, []string{}, name, property)
}

func PropertySchema(r *cfschema.Resource, pathPrefix []string, name string, property *cfschema.Property) (string, Features) {
	var features Features

	if r == nil {
		return "", features
	}

	var b strings.Builder
	path := append(pathPrefix, name)
	indentation := strings.Repeat("\t", len(path))

	createOnly := r.CreateOnlyProperties.ContainsPath(path)
	readOnly := r.ReadOnlyProperties.ContainsPath(path)
	required := r.IsRequired(name)

	if name != "" {
		fmt.Fprintf(&b, "\n\n%s\"%s\": {", indentation, strcase.ToSnake(name))
	}

	switch property.Type.String() {
	default:
		fmt.Fprintf(&b, "\n%sType: Unknown,", indentation)
	case cfschema.PropertyTypeArray:
		if property.InsertionOrder != nil && *property.InsertionOrder {
			fmt.Fprintf(&b, "\n%sType: schema.TypeList,", indentation)
		} else {
			fmt.Fprintf(&b, "\n%sType: schemaTypeSet,", indentation)
		}

		if property.MaxItems != nil {
			fmt.Fprintf(&b, "\n%sMaxItems: %d,", indentation, *property.MaxItems)
		}

		if property.MinItems != nil {
			fmt.Fprintf(&b, "\n%sMinItems: %d,", indentation, *property.MinItems)
		}

		if property.Items != nil {
			if property.Items.Type.String() == cfschema.PropertyTypeObject {
				fmt.Fprintf(&b, "\n%sElem: &schema.Resource{", indentation)
				fmt.Fprintf(&b, "\n%s\tSchema: map[string]schema.Schema{", indentation)
			} else {
				fmt.Fprintf(&b, "\n%sElem: &schema.Schema{", indentation)
			}

			ps, f := PropertySchema(r, path, "", property.Items)
			fmt.Fprint(&b, ps)
			features |= f

			if property.Items.Type.String() == cfschema.PropertyTypeObject {
				fmt.Fprintf(&b, "\n%s\t},", indentation)
			}

			fmt.Fprintf(&b, "\n%s},", indentation)
		}
	case cfschema.PropertyTypeBoolean:
		fmt.Fprintf(&b, "\n%sType: schema.TypeBool,", indentation)
	case cfschema.PropertyTypeInteger:
		fmt.Fprintf(&b, "\n%sType: schema.TypeInt,", indentation)

		if len(property.Enum) == 0 {
			break
		}

		fmt.Fprintf(&b, "\n%sValidateFunc: validation.All(", indentation)
		features |= UsesValidation

		if len(property.Enum) > 0 {
			fmt.Fprintf(&b, "\n%s\tvalidation.IntInSlice([]int{", indentation)
			for _, enumItem := range property.Enum {
				switch v := enumItem.(type) {
				case float64:
					fmt.Fprintf(&b, "\n%s\t\t%d,", indentation, int(v))
				case int64:
					fmt.Fprintf(&b, "\n%s\t\t%d,", indentation, v)
				}
			}

			fmt.Fprintf(&b, "\n%s\t}),", indentation)
		}

		fmt.Fprintf(&b, "\n%s),", indentation)
	case cfschema.PropertyTypeNumber:
		fmt.Fprintf(&b, "\n%sType: schema.TypeFloat,", indentation)
	case cfschema.PropertyTypeObject:
		// If there are no underlying Properties, the schema is not defined.
		// CloudFormation documentation denotes these as Json or Object.
		if len(property.Properties) == 0 {
			fmt.Fprintf(&b, "\n%sType: TypeDynamic,", indentation)
			break
		}

		fmt.Fprintf(&b, "\n%sType: schema.TypeList,", indentation)
		fmt.Fprintf(&b, "\n%sMaxItems: 1,", indentation)
		fmt.Fprintf(&b, "\n%sElem: &schema.Resource{", indentation)
		fmt.Fprintf(&b, "\n%s\tSchema: map[string]schema.Schema{", indentation)
		for objPropertyName, objProperty := range property.Properties {
			ps, f := PropertySchema(r, path, objPropertyName, objProperty)
			fmt.Fprint(&b, ps)
			features |= f
		}
		fmt.Fprintf(&b, "\n%s\t},", indentation)
		fmt.Fprintf(&b, "\n%s},", indentation)
	case cfschema.PropertyTypeString:
		fmt.Fprintf(&b, "\n%sType: schema.TypeString,", indentation)

		if len(property.Enum) == 0 && property.MaxLength == nil && property.MinLength == nil && property.Pattern == nil {
			break
		}

		fmt.Fprintf(&b, "\n%sValidateFunc: validation.All(", indentation)
		features |= UsesValidation

		if len(property.Enum) > 0 {
			fmt.Fprintf(&b, "\n%s\tvalidation.StringInSlice([]string{", indentation)
			for _, enumItem := range property.Enum {
				fmt.Fprintf(&b, "\n%s\t\t%q,", indentation, enumItem)
			}
			fmt.Fprintf(&b, "\n%s\t}, false),", indentation)
		}

		if property.MaxLength != nil && property.MinLength != nil {
			fmt.Fprintf(&b, "\n%s\tvalidation.StringLenBetween(%d, %d),", indentation, *property.MinLength, *property.MaxLength)
		} else if property.MaxLength != nil {
			fmt.Fprintf(&b, "\n%s\tvalidation.StringLenBetween(0, %d),", indentation, *property.MaxLength)
		}

		if property.Pattern != nil {
			fmt.Fprintf(&b, "\n%s\tvalidation.StringMatch(regexp.MustCompile(`%s`), \"\"),", indentation, *property.Pattern)
			features |= UsesRegexp
		}

		fmt.Fprintf(&b, "\n%s),", indentation)
	}

	// Array items
	if name == "" {
		fmt.Fprintf(&b, "\n%s},", indentation)
		return b.String(), features
	}

	if required {
		fmt.Fprintf(&b, "\n%sRequired: true,", indentation)
	} else if !readOnly {
		fmt.Fprintf(&b, "\n%sOptional: true,", indentation)
	}

	if readOnly && !required {
		fmt.Fprintf(&b, "\n%sComputed: true,", indentation)
	}

	if createOnly {
		fmt.Fprintf(&b, "\n%sForceNew: true,", indentation)
	}

	if !createOnly && !readOnly {
		features |= HasUpdatableProperty
	}

	if property.Description != nil {
		fmt.Fprintf(&b, "\n%sDescription: `%s`,", indentation, *property.Description)
	}

	fmt.Fprintf(&b, "\n%s},", indentation)

	return b.String(), features
}