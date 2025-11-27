package util

type ISCMapToMap[K comparable, V any, R any] struct {
	BsMap[K, V]
}

func MapToMapFrom[K comparable, V any, R any](m BsMap[K, V]) ISCMapToMap[K, V, R] {
	return ISCMapToMap[K, V, R]{
		m,
	}
}

func (m ISCMapToMap[K, V, R]) FlatMap(f func(K, V) []R) BsList[R] {
	return MapFlatMap(m.BsMap, f)
}

func (m ISCMapToMap[K, V, R]) FlatMapTo(dest *[]R, f func(K, V) []R) BsList[R] {
	return MapFlatMapTo(m.BsMap, dest, f)
}

func (m ISCMapToMap[K, V, R]) Map(f func(K, V) R) BsList[R] {
	return MapMap(m.BsMap, f)
}

func (m ISCMapToMap[K, V, R]) MapTo(dest *[]R, f func(K, V) R) BsList[R] {
	return MapMapTo(m.BsMap, dest, f)
}
