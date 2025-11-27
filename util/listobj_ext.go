package util

type ISCListToMap[T any, R any] struct {
	BsList[T]
}

func ListToMapFrom[T any, R any](list BsList[T]) ISCListToMap[T, R] {
	return ISCListToMap[T, R]{
		list,
	}
}

func (l ISCListToMap[T, R]) FlatMap(f func(T) []R) BsList[R] {
	return ListFlatMap(l.BsList, f)
}

func (l ISCListToMap[T, R]) FlatMapIndexed(f func(int, T) []R) BsList[R] {
	return ListFlatMapIndexed(l.BsList, f)
}

func (l ISCListToMap[T, R]) FlatMapTo(dest *[]R, f func(T) []R) BsList[R] {
	return ListFlatMapTo(l.BsList, dest, f)
}

func (l ISCListToMap[T, R]) FlatMapIndexedTo(dest *[]R, f func(int, T) []R) BsList[R] {
	return ListFlatMapIndexedTo(l.BsList, dest, f)
}

func (l ISCListToMap[T, R]) Map(f func(T) R) BsList[R] {
	return ListMap(l.BsList, f)
}

func (l ISCListToMap[T, R]) MapIndexed(f func(int, T) R) BsList[R] {
	return ListMapIndexed(l.BsList, f)
}

func (l ISCListToMap[T, R]) MapTo(dest *[]R, f func(T) R) BsList[R] {
	return ListMapTo(l.BsList, dest, f)
}

func (l ISCListToMap[T, R]) MapIndexedTo(dest *[]R, f func(int, T) R) BsList[R] {
	return ListMapIndexedTo(l.BsList, dest, f)
}

func (l ISCListToMap[T, R]) Reduce(init func(T) R, f func(R, T) R) R {
	return Reduce(l.BsList, init, f)
}

func (l ISCListToMap[T, R]) ReduceIndexed(init func(int, T) R, f func(int, R, T) R) R {
	return ReduceIndexed(l.BsList, init, f)
}

type ISCListToSlice[T any, R comparable] struct {
	BsList[T]
}

func ListToSliceFrom[T any, R comparable](list BsList[T]) ISCListToSlice[T, R] {
	return ISCListToSlice[T, R]{
		list,
	}
}

func (l ISCListToSlice[T, R]) SliceContains(predicate func(T) R, key R) bool {
	return SliceContains(l.BsList, predicate, key)
}

func (l ISCListToSlice[T, R]) SliceTo(valueTransform func(T) R) BsMap[R, T] {
	return SliceTo(l.BsList, valueTransform)
}

type ISCListToTriple[T comparable, K comparable, V any] struct {
	BsList[T]
}

func ListToTripleFrom[T comparable, K comparable, V any](list BsList[T]) ISCListToTriple[T, K, V] {
	return ISCListToTriple[T, K, V]{
		list,
	}
}

func (l ISCListToTriple[T, K, V]) GroupBy(f func(T) K) map[K][]T {
	return GroupBy(l.BsList, f)
}

func (l ISCListToTriple[T, K, V]) GroupByTransform(f func(T) K, trans func(T) V) map[K][]V {
	return GroupByTransform(l.BsList, f, trans)
}

func (l ISCListToTriple[T, K, V]) GroupByTo(dest *map[K][]T, f func(T) K) map[K][]T {
	return GroupByTo(l.BsList, dest, f)
}

func (l ISCListToTriple[T, K, V]) GroupByTransformTo(dest *map[K][]V, f func(T) K, trans func(T) V) map[K][]V {
	return GroupByTransformTo(l.BsList, dest, f, trans)
}

func (l ISCListToTriple[T, K, V]) Associate(transform func(T) Pair[K, V]) BsMap[K, V] {
	return Associate(l.BsList, transform)
}

func (l ISCListToTriple[T, K, V]) AssociateTo(destination *map[K]V, transform func(T) Pair[K, V]) BsMap[K, V] {
	return AssociateTo(l.BsList, destination, transform)
}

func (l ISCListToTriple[T, K, V]) AssociateBy(keySelector func(T) K) BsMap[K, T] {
	return AssociateBy(l.BsList, keySelector)
}

func (l ISCListToTriple[T, K, V]) AssociateByAndValue(keySelector func(T) K, valueTransform func(T) V) BsMap[K, V] {
	return AssociateByAndValue(l.BsList, keySelector, valueTransform)
}

func (l ISCListToTriple[T, K, V]) AssociateByTo(destination *map[K]T, keySelector func(T) K) BsMap[K, T] {
	return AssociateByTo(l.BsList, destination, keySelector)
}

func (l ISCListToTriple[T, K, V]) AssociateByAndValueTo(destination *map[K]V, keySelector func(T) K, valueTransform func(T) V) BsMap[K, V] {
	return AssociateByAndValueTo(l.BsList, destination, keySelector, valueTransform)
}

func (l ISCListToTriple[T, K, V]) AssociateWith(valueSelector func(T) V) BsMap[T, V] {
	return AssociateWith(l.BsList, valueSelector)
}

func (l ISCListToTriple[T, K, V]) AssociateWithTo(destination *map[T]V, valueSelector func(T) V) BsMap[T, V] {
	return AssociateWithTo(l.BsList, destination, valueSelector)
}

type ISCListToPair[K comparable, V comparable] struct {
	BsList[Pair[K, V]]
}

func ListToPairFrom[K comparable, V comparable](list BsList[Pair[K, V]]) ISCListToPair[K, V] {
	return ISCListToPair[K, V]{
		list,
	}
}

func ListToPairWithPairs[K comparable, V comparable](list ...Pair[K, V]) ISCListToPair[K, V] {
	return ISCListToPair[K, V]{
		list,
	}
}

func (l ISCListToPair[K, V]) ToMap() BsMap[K, V] {
	m := make(map[K]V)
	for _, item := range l.BsList {
		m[item.First] = item.Second
	}
	return NewMapWithMap(m)
}
