package config

import (
	"fmt"

	"github.com/simonalong/gole/listener"
)

var EventOfConfigChange = "event_of_config_change"
var EventOfConfigLoadFinish = "event_of_config_load_finish"

type EventOfLoadFinish struct{}

func (e EventOfLoadFinish) Name() string {
	return EventOfConfigLoadFinish
}

func (e EventOfLoadFinish) Group() string { return listener.DefaultGroup }

func (e EventOfLoadFinish) ToString() string {
	return fmt.Sprintf("%v", listener.DefaultGroup)
}

// EventOfChange 配置变更事件, 对应：event_of_config_change
type EventOfChange struct {
	GroupName string
	Key       string
	Value     string
}

func (e EventOfChange) Name() string { return EventOfConfigChange }

func (e EventOfChange) Group() string {
	return e.GroupName
}

func (e EventOfChange) ToString() string {
	return fmt.Sprintf("%v:%v=%v", e.GroupName, e.Key, e.Value)
}
