package global

import (
	"context"
	"github.com/simonalong/gole/goid"
)
import "go.opentelemetry.io/otel/trace"

var Tracer trace.Tracer
var ContextLocalStorage goid.LocalStorage

func GetGlobalContext() context.Context {
	if ContextLocalStorage != nil {
		ctx := ContextLocalStorage.Get()
		if ctx != nil {
			return ctx.(context.Context)
		}
	}
	return context.Background()
}

func SetGlobalContext(ctx context.Context) {
	if ContextLocalStorage != nil {
		ContextLocalStorage.Set(ctx)
	}
}
