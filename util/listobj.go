package util

type BsList[T any] []T

func NewList[T any]() BsList[T] {
	return []T{}
}

func NewListWithList[T any](list []T) BsList[T] {
	return list
}

func NewListWithItems[T any](items ...T) BsList[T] {
	return items
}

func (l *BsList[T]) Add(item T) int {
	idx := len(*l)
	*l = append(*l, item)
	return idx
}

func (l *BsList[T]) AddAll(item ...T) {
	*l = append(*l, item...)
}

func (l *BsList[T]) Insert(index int, item T) int {
	*l = append((*l)[:index], append([]T{item}, (*l)[index:]...)...)
	return index
}

func (l *BsList[T]) Delete(index int) T {
	item := (*l)[index]
	*l = append((*l)[:index], (*l)[index+1:]...)
	return item
}

func (l *BsList[T]) Clear() {
	*l = []T{}
}

func (l BsList[T]) IsEmpty() bool {
	return len(l) == 0
}

func (l BsList[T]) Size() int {
	return len(l)
}

func (l BsList[T]) ForEach(f func(T)) {
	for _, item := range l {
		f(item)
	}
}

func (l BsList[T]) ForEachIndexed(f func(int, T)) {
	for idx, item := range l {
		f(idx, item)
	}
}

func (l BsList[T]) Distinct() BsList[T] {
	return ListDistinct(l)
}

func (l BsList[T]) Filter(f func(T) bool) BsList[T] {
	return ListFilter(l, f)
}

func (l BsList[T]) FilterNot(f func(T) bool) BsList[T] {
	return ListFilterNot(l, f)
}

func (l BsList[T]) FilterIndexed(f func(int, T) bool) BsList[T] {
	return ListFilterIndexed(l, f)
}

func (l BsList[T]) FilterNotIndexed(f func(int, T) bool) BsList[T] {
	return ListFilterNotIndexed(l, f)
}

func (l BsList[T]) FilterTo(dest *[]T, f func(T) bool) BsList[T] {
	return ListFilterTo(l, dest, f)
}

func (l BsList[T]) FilterNotTo(dest *[]T, f func(T) bool) BsList[T] {
	return ListFilterNotTo(l, dest, f)
}

func (l BsList[T]) FilterIndexedTo(dest *[]T, f func(int, T) bool) BsList[T] {
	return ListFilterIndexedTo(l, dest, f)
}

func (l BsList[T]) FilterNotIndexedTo(dest *[]T, f func(int, T) bool) BsList[T] {
	return ListFilterNotIndexedTo(l, dest, f)
}

func (l BsList[T]) Contains(item T) bool {
	return ListContains(l, item)
}

func (l BsList[T]) Find(f func(T) bool) *T {
	return Find(l, f)
}

func (l BsList[T]) FindLast(f func(T) bool) *T {
	return FindLast(l, f)
}

func (l BsList[T]) First() T {
	return First(l)
}

func (l BsList[T]) Last() T {
	return Last(l)
}

func (l BsList[T]) FirstOrNull() *T {
	return FirstOrNull(l)
}

func (l BsList[T]) LastOrNull() *T {
	return LastOrNull(l)
}

func (l BsList[T]) IndexOf(item T) int {
	return IndexOf(l, item)
}

func (l BsList[T]) LastIndexOf(item T) int {
	return LastIndexOf(l, item)
}

func (l BsList[T]) IndexOfCondition(f func(T) bool) int {
	return IndexOfCondition(l, f)
}

func (l BsList[T]) LastIndexOfCondition(f func(T) bool) int {
	return LastIndexOfCondition(l, f)
}

func (l BsList[T]) JoinToString(f func(T) string) string {
	return ListJoinToString(l, f)
}

func (l BsList[T]) JoinToStringFull(sep string, prefix string, postfix string, f func(T) string) string {
	return ListJoinToStringFull(l, sep, prefix, postfix, f)
}

func (l BsList[T]) All(f func(T) bool) bool {
	return ListAll(l, f)
}

func (l BsList[T]) Any(f func(T) bool) bool {
	return ListAny(l, f)
}

func (l BsList[T]) None(f func(T) bool) bool {
	return ListNone(l, f)
}

func (l BsList[T]) Count(f func(T) bool) int {
	return ListCount(l, f)
}

func (l BsList[T]) SubList(fromIndex int, toIndex int) BsList[T] {
	r := SubList(l, fromIndex, toIndex)
	return NewListWithList(r)
}

func (l BsList[T]) Slice(r IntRange) BsList[T] {
	rr := Slice(l, r)
	return NewListWithList(rr)
}

func (l BsList[T]) Take(n int) BsList[T] {
	r := Take(l, n)
	return NewListWithList(r)
}

func (l BsList[T]) TakeLast(n int) BsList[T] {
	r := TakeLast(l, n)
	return NewListWithList(r)
}

func (l BsList[T]) TakeWhile(n int, f func(T) bool) BsList[T] {
	r := TakeWhile(l, n, f)
	return NewListWithList(r)
}

func (l BsList[T]) TakeLastWhile(n int, f func(T) bool) BsList[T] {
	r := TakeLastWhile(l, n, f)
	return NewListWithList(r)
}

func (l BsList[T]) Drop(n int) BsList[T] {
	r := Drop(l, n)
	return NewListWithList(r)
}

func (l BsList[T]) DropLast(n int) BsList[T] {
	r := DropLast(l, n)
	return NewListWithList(r)
}

func (l BsList[T]) DropWhile(n int, f func(T) bool) BsList[T] {
	r := DropWhile(l, n, f)
	return NewListWithList(r)
}

func (l BsList[T]) DropLastWhile(n int, f func(T) bool) BsList[T] {
	r := DropLastWhile(l, n, f)
	return NewListWithList(r)
}

func (l BsList[T]) Partition(partition int) [][]T {
	return Partition(l, partition)
}

func (l BsList[T]) PartitionWithCal(f func(int) int) [][]T {
	return PartitionWithCal(l, f)
}

func (l BsList[T]) Plus(n []T) BsList[T] {
	return ListPlus(l, n)
}

func (l BsList[T]) Minus(n []T) BsList[T] {
	return ListMinus(l, n)
}

func (l BsList[T]) Equals(n BsList[T]) bool {
	return ListEquals(l, n)
}

func ListToSet[T comparable](list BsList[T]) Set[T] {
	res := Set[T]{}
	for _, v := range list {
		_ = res.Add(v)
	}
	return res
}
