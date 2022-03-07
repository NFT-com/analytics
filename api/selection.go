package api

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/graphql"
)

func getSelections(ctx context.Context) []string {

	opctx := graphql.GetOperationContext(ctx)
	fields := graphql.CollectFieldsCtx(ctx, nil)

	return getNestedSelections(
		opctx,
		fields,
		"",
	)
}

func getNestedSelections(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) []string {

	var selections []string

	for _, field := range fields {

		prefixed := prefixField(prefix, field.Name)

		log.Printf("prefixed field: %v", prefixed)
		selections = append(selections, prefixed)

		nestedSelections := getNestedSelections(ctx, graphql.CollectFields(ctx, field.Selections, nil), prefixed)
		selections = append(selections, nestedSelections...)
	}

	return selections
}

// return name or prefix.name
func prefixField(prefix string, name string) string {
	if prefix == "" {
		return name
	}

	return prefix + "." + name
}
