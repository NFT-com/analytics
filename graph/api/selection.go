package api

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// querySelection represents the selection set for a query.
// Top-level fields will be shown by name, while nested fields will be shown as `parent.field`.
type querySelection struct {
	fields map[string]struct{}
}

func (q querySelection) isSelected(name string) bool {
	_, ok := q.fields[name]
	return ok
}

func getQuerySelection(ctx context.Context) querySelection {

	op := graphql.GetOperationContext(ctx)
	collected := graphql.CollectFieldsCtx(ctx, nil)

	// Get list of selected fields.
	fields := getNestedSelection(op, collected, "")

	// Transform list of selected fields into a map for easier lookup.
	fieldMap := make(map[string]struct{})
	for _, field := range fields {
		fieldMap[field] = struct{}{}
	}

	qs := querySelection{
		fields: fieldMap,
	}

	return qs
}

func getNestedSelection(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) []string {

	requested := make([]string, 0)

	for _, field := range fields {
		name := formatField(prefix, field.Name)
		requested = append(requested, name)

		collected := graphql.CollectFields(ctx, field.Selections, nil)
		requested = append(requested, getNestedSelection(ctx, collected, name)...)
	}

	return requested
}

func formatField(fields ...string) string {

	if len(fields) == 0 {
		return ""
	}

	if fields[0] == "" {
		fields = fields[1:]
	}

	out := fields[0]
	for _, field := range fields[1:] {
		out += "." + field
	}

	return out
}
