package util

import (
	"fmt"
)

type Set[T comparable] map[T]struct{}

// NewSet 初始化并指定存储对象的类型
func NewSet[T comparable]() Set[T] {
	return Set[T]{}
}

func NewSetWithList[T comparable](list []T) Set[T] {
	s := Set[T]{}
	s.AddAll(list...)
	return s
}

func NewSetWithItems[T comparable](items ...T) Set[T] {
	s := Set[T]{}
	s.AddAll(items...)
	return s
}

// Size 返回数据数量
func (s Set[T]) Size() int {
	return len(s)
}

// Add 添加元素
func (s *Set[T]) Add(item T) error {
	if !s.Contains(item) {
		(*s)[item] = struct{}{}
		return nil
	} else {
		return fmt.Errorf("%v already exists in set", item)
	}
}

// AddAll 添加多个元素
func (s *Set[T]) AddAll(items ...T) {
	for _, item := range items {
		_ = s.Add(item)
	}
}

// Delete 删除指定Key元素
func (s *Set[T]) Delete(item T) error {
	if s.Contains(item) {
		delete(*s, item)
		return nil
	} else {
		return fmt.Errorf("%v not exists in set", item)
	}
}

// Contains 判断key是否存在
func (s Set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

// Clear 重置
func (s *Set[T]) Clear() {
	*s = Set[T]{}
}

func (s Set[T]) ToList() BsList[T] {
	res := NewList[T]()
	for k := range s {
		res.Add(k)
	}
	return res
}
