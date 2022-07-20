package query

import (
	"context"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)

// Selection represents the selection set for a query. Top-level fields are
// shown by name, while nested fields are shown as `parent.field`.
type Selection struct {
	fields map[string]Arguments
}

// Arguments represents the arguments given for a specific GraphQL field path.
type Arguments map[string]interface{}

type queryField struct {
	path      string
	arguments Arguments
}

// GetSelection returns the selection set from a query. The provided context
// MUST be a context provided by the `gqlgen` framework, containing an operation
// context.
func GetSelection(ctx context.Context) Selection {

	op := graphql.GetOperationContext(ctx)
	collected := graphql.CollectFieldsCtx(ctx, nil)

	// Get list of selected fields.
	fields := getNestedSelection(op, collected, "")

	// Transform list of selected fields into a map for easier lookup.
	fieldMap := make(map[string]Arguments)
	for _, field := range fields {
		fieldMap[field.path] = field.arguments
	}

	s := Selection{
		fields: fieldMap,
	}

	return s
}

// Has returns true if the specified field is found in the selection set.
func (s *Selection) Has(name string) bool {
	_, ok := s.fields[name]
	return ok
}

// Args returns the arguments for the specified fields.
func (s *Selection) Args(name string) Arguments {
	return s.fields[name]
}

// getNestedSelection returns the selected fields, prefixed with their parent path.
func getNestedSelection(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) []queryField {

	var requested []queryField

	for _, field := range fields {

		// Get the formatted name for the field.
		var name string
		if prefix != "" {
			name = FieldPath(prefix, field.Name)
		} else {
			name = FieldPath(field.Name)
		}

		// Get all arguments for the field.
		m := make(map[string]interface{})
		args := field.ArgumentMap(m)

		qf := queryField{
			path:      name,
			arguments: args,
		}

		requested = append(requested, qf)

		collected := graphql.CollectFields(ctx, field.Selections, nil)
		requested = append(requested, getNestedSelection(ctx, collected, name)...)
	}

	return requested
}

// FieldPath returns the selection path for the field, based on provided path components.
func FieldPath(fields ...string) string {

	path := strings.Join(fields, ".")
	return path
}
