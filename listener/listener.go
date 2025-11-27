package listener

import cmap "github.com/orcaman/concurrent-map"

type EventListener func(event BaseEvent)

var eventWatcherGroupMaps cmap.ConcurrentMap

func init() {
	eventWatcherGroupMaps = cmap.New()
}

func PublishEvent(event BaseEvent) {
	if eventWatcherGroupMaps.Has(event.Group()) {
		_eventWatcherGroup, _ := eventWatcherGroupMaps.Get(event.Group())
		eventWatcherGroup := _eventWatcherGroup.(cmap.ConcurrentMap)
		if eventWatchers, exist := eventWatcherGroup.Get(event.Name()); exist {
			for _, eventWatcher := range eventWatchers.([]EventListener) {
				eventWatcher(event)
			}
		}
	}

	// 选择监听所有分组的事件
	if eventWatcherGroupMaps.Has("*") {
		_eventWatcherGroup, _ := eventWatcherGroupMaps.Get("*")
		eventWatcherGroup := _eventWatcherGroup.(cmap.ConcurrentMap)
		if eventWatchers, exist := eventWatcherGroup.Get(event.Name()); exist {
			for _, eventWatcher := range eventWatchers.([]EventListener) {
				eventWatcher(event)
			}
		}
	}
}

func AddListener(eventName string, eventListener EventListener) {
	AddListenerWithGroup(DefaultGroup, eventName, eventListener)
}

// AddListenerWithGroup 增加事件监听器
// group 分组，eventName 事件名称，eventListener 事件监听器
// group 也支持：* 表示监听所有分组的事件
func AddListenerWithGroup(group string, eventName string, eventListener EventListener) {
	if eventWatcherGroupMaps.Has(group) {
		_eventWatcherMap, _ := eventWatcherGroupMaps.Get(group)
		eventWatcherMap := _eventWatcherMap.(cmap.ConcurrentMap)
		if eventWatchers, exist := eventWatcherMap.Get(eventName); exist {
			eventWatchers = append(eventWatchers.([]EventListener), eventListener)
			eventWatcherMap.Set(eventName, eventWatchers)
		} else {
			eventWatchers = []EventListener{}
			eventWatchers = append(eventWatchers.([]EventListener), eventListener)
			eventWatcherMap.Set(eventName, eventWatchers)
		}
	} else {
		var eventWatchers []EventListener
		eventWatchers = append(eventWatchers, eventListener)

		eventMap := cmap.New()
		eventMap.Set(eventName, eventWatchers)
		eventWatcherGroupMaps.Set(group, eventMap)
	}
}

func ContainListener(group, eventName string) bool {
	if eventWatcherGroupMaps.Has(group) {
		_eventWatcherMap, _ := eventWatcherGroupMaps.Get(group)
		eventWatcherMap := _eventWatcherMap.(cmap.ConcurrentMap)
		if _, exist := eventWatcherMap.Get(eventName); exist {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
