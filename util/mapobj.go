package util

type BsMap[K comparable, V any] map[K]V

func NewMap[K comparable, V any]() BsMap[K, V] {
	return make(BsMap[K, V])
}

func NewMapWithMap[K comparable, V any](ma map[K]V) BsMap[K, V] {
	return ma
}

func NewMapWithPairs[K comparable, V any](pairs ...Pair[K, V]) BsMap[K, V] {
	m := make(map[K]V)
	for _, item := range pairs {
		m[item.First] = item.Second
	}
	return m
}

func (m BsMap[K, V]) Size() int {
	return len(m)
}

func (m BsMap[K, V]) Put(k K, v V) {
	m[k] = v
}

func (m BsMap[K, V]) PutPair(item Pair[K, V]) {
	m[item.First] = item.Second
}

func (m BsMap[K, V]) PutAllPairs(item ...Pair[K, V]) {
	for _, e := range item {
		m.PutPair(e)
	}
}

func (m BsMap[K, V]) Get(k K) V {
	return m[k]
}

func (m BsMap[K, V]) GetOrDef(k K, def V) V {
	if v, ok := m[k]; ok {
		return v
	} else {
		return def
	}
}

func (m BsMap[K, V]) Delete(k K) {
	delete(m, k)
}

func (m *BsMap[K, V]) Clear() {
	*m = make(BsMap[K, V])
}

func (m BsMap[K, V]) ForEach(f func(K, V)) {
	for k, v := range m {
		f(k, v)
	}
}

func (m BsMap[K, V]) Filter(f func(K, V) bool) BsMap[K, V] {
	return MapFilter(m, f)
}

func (m BsMap[K, V]) FilterNot(f func(K, V) bool) BsMap[K, V] {
	return MapFilterNot(m, f)
}

func (m BsMap[K, V]) FilterKeys(f func(K) bool) BsMap[K, V] {
	return MapFilterKeys(m, f)
}

func (m BsMap[K, V]) FilterValues(f func(V) bool) BsMap[K, V] {
	return MapFilterValues(m, f)
}

func (m BsMap[K, V]) FilterTo(dest *map[K]V, f func(K, V) bool) BsMap[K, V] {
	return MapFilterTo(m, dest, f)
}

func (m BsMap[K, V]) FilterNotTo(dest *map[K]V, f func(K, V) bool) BsMap[K, V] {
	return MapFilterNotTo(m, dest, f)
}

func (m BsMap[K, V]) Contains(k K, v V) bool {
	return MapContains(m, k, v)
}

func (m BsMap[K, V]) ContainsKey(k K) bool {
	return MapContainsKey(m, k)
}

func (m BsMap[K, V]) ContainsValue(v V) bool {
	return MapContainsValue(m, v)
}

func (m BsMap[K, V]) JoinToString(f func(K, V) string) string {
	return MapJoinToString(m, f)
}

func (m BsMap[K, V]) JoinToStringFull(sep string, prefix string, postfix string, f func(K, V) string) string {
	return MapJoinToStringFull(m, sep, prefix, postfix, f)
}

func (m BsMap[K, V]) All(f func(K, V) bool) bool {
	return MapAll(m, f)
}

func (m BsMap[K, V]) Any(f func(K, V) bool) bool {
	return MapAny(m, f)
}

func (m BsMap[K, V]) None(f func(K, V) bool) bool {
	return MapNone(m, f)
}

func (m BsMap[K, V]) Count(f func(K, V) bool) int {
	return MapCount(m, f)
}

func (m BsMap[K, V]) AllKey(f func(K) bool) bool {
	return MapAllKey(m, f)
}

func (m BsMap[K, V]) AnyKey(f func(K) bool) bool {
	return MapAnyKey(m, f)
}

func (m BsMap[K, V]) NoneKey(f func(K) bool) bool {
	return MapNoneKey(m, f)
}

func (m BsMap[K, V]) CountKey(f func(K) bool) int {
	return MapCountKey(m, f)
}

func (m BsMap[K, V]) AllValue(f func(V) bool) bool {
	return MapAllValue(m, f)
}

func (m BsMap[K, V]) AnyValue(f func(V) bool) bool {
	return MapAnyValue(m, f)
}

func (m BsMap[K, V]) NoneValue(f func(V) bool) bool {
	return MapNoneValue(m, f)
}

func (m BsMap[K, V]) CountValue(f func(V) bool) int {
	return MapCountValue(m, f)
}

func (m BsMap[K, V]) ToList() []Pair[K, V] {
	var n []Pair[K, V]
	for k, v := range m {
		n = append(n, NewPair(k, v))
	}
	return n
}

func (m BsMap[K, V]) Plus(n map[K]V) BsMap[K, V] {
	return MapPlus(m, n)
}

func (m BsMap[K, V]) Minus(n map[K]V) BsMap[K, V] {
	return MapMinus(m, n)
}

func (m BsMap[K, V]) Equals(n map[K]V) bool {
	return MapEquals(m, n)
}

func (m BsMap[K, V]) Keys() BsList[K] {
	i := 0
	keys := make([]K, len(m))
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
