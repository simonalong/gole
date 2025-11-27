package listener

var DefaultGroup = "default"

type BaseEvent interface {
	Name() string
	Group() string
	ToString() string
}
