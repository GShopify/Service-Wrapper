package server

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/opentracing/opentracing-go/log"
	"github.com/vektah/gqlparser/v2/ast"
)

func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

func GetNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}

func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}

	return name
}

func GetFieldValue(ctx context.Context, field string, arg string) any {
	var (
		fields = graphql.GetFieldContext(ctx).Field.Selections
		f      *ast.Field
		ok     bool
	)

	for k := range fields {
		if f, ok = fields[k].(*ast.Field); !ok || f.Name != field {
			continue
		}

		if key := f.Arguments.ForName(arg); key != nil {
			val, err := key.Value.Value(graphql.GetOperationContext(ctx).Variables)
			if err != nil {
				log.Error(err)
			}

			return val
		}
	}

	return nil
}
