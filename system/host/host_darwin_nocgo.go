//go:build darwin && !cgo

package host

import (
	"context"
	"github.com/simonalong/gole/system/common"
)

func SensorsTemperaturesWithContext(ctx context.Context) ([]TemperatureStat, error) {
	return []TemperatureStat{}, common.ErrNotImplementedError
}
