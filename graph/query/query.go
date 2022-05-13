package query

import (
	"context"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)

// Selection represents the selection set for a query. Top-level fields are
// shown by name, while nested fields are shown as `parent.field`.
type Selection struct {
	fields map[string]struct{}
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
	fieldMap := make(map[string]struct{})
	for _, field := range fields {
		fieldMap[field] = struct{}{}
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

// getNestedSelection returns the selected fields, prefixed with their parent path.
func getNestedSelection(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) []string {

	requested := make([]string, 0)

	for _, field := range fields {
		name := FieldPath(prefix, field.Name)
		requested = append(requested, name)

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
