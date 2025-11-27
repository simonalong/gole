package global

import (
	"github.com/simonalong/gole/config"
	"go.opentelemetry.io/otel/trace"
)

func GetTraceId() string {
	if config.GetValueBoolDefault("gole.opentelemetry.enable", true) {
		traceId := trace.SpanFromContext(GetGlobalContext()).SpanContext().TraceID().String()
		if traceId == "00000000000000000000000000000000" {
			traceId = ""
		}
		return traceId
	}
	return ""
}

func SetTraceId(traceId string) error {
	tid, err := trace.TraceIDFromHex(traceId)
	if err != nil {
		return err
	}
	newSc := trace.SpanFromContext(GetGlobalContext()).SpanContext().WithTraceID(tid)
	ctx := trace.ContextWithRemoteSpanContext(GetGlobalContext(), newSc)
	SetGlobalContext(ctx)
	return err
}
