package util

import (
	"sync"
)

type inter interface{}

type Set struct {
	m map[inter]bool
	sync.RWMutex
}

func NewSet() *Set {
	return &Set{
		m: map[inter]bool{},
	}
}

func (s *Set) Add(itme inter) {
	s.Lock()
	defer s.Unlock()
	s.m[itme] = true
}

func (s *Set) Remove(item inter) {
	s.Lock()
	delete(s.m, item)
	s.Unlock()
}

func (s *Set) Has(item inter) bool {
	s.Lock()
	_, ok := s.m[item]
	s.Unlock()
	return ok
}

func (s *Set) Len() int {
	return len(s.List())
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[inter]bool{}
}

func (s *Set) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Set) List() []inter {
	s.RLock()
	defer s.RUnlock()
	list := []inter{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
