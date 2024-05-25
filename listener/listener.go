package listener

type EventListener func(event GoleEvent)

var eventWatcherGroupMaps map[string]map[string][]EventListener

type GoleEvent interface {
	Name() string
	Group() string
}

func PublishEvent(event GoleEvent) {
	if eventWatcherGroupMaps == nil {
		return
	}
	if eventWatcherGroup, exist := eventWatcherGroupMaps[event.Group()]; exist {
		if eventWatchers, exist := eventWatcherGroup[event.Name()]; exist {
			for _, eventWatcher := range eventWatchers {
				eventWatcher(event)
			}
		}
	}
}

func AddListener(eventName string, eventListener EventListener) {
	AddListenerWithGroup(DefaultGroup, eventName, eventListener)
}

func AddListenerWithGroup(group string, eventName string, eventListener EventListener) {
	if eventWatcherGroupMaps == nil {
		eventWatcherGroupMaps = map[string]map[string][]EventListener{}
	}

	if eventWatcherMap, exist := eventWatcherGroupMaps[group]; exist {
		if eventWatchers, exist := eventWatcherMap[eventName]; exist {
			eventWatchers = append(eventWatchers, eventListener)
			eventWatcherMap[eventName] = eventWatchers
		} else {
			eventWatchers = []EventListener{}
			eventWatchers = append(eventWatchers, eventListener)
			eventWatcherMap[eventName] = eventWatchers
		}
	} else {
		eventWatchers := []EventListener{}
		eventWatchers = append(eventWatchers, eventListener)

		eventMap := map[string][]EventListener{}
		eventMap[eventName] = eventWatchers
		eventWatcherGroupMaps[group] = eventMap
	}
}
