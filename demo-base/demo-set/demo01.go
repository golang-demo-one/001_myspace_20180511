package main

import (
	"fmt"
	"sync"
)

/*
	https://studygolang.com/articles/2335
*/

type ISet interface {
	Add(item int)
	Remove(item int)
	Has(item int)
	Len() int
	IsEmpty()
	list()
	Clear()
}

type Set struct {
	m map[int]bool
	sync.RWMutex
}

func New() *Set {
	return &Set{
		m: map[int]bool{},
	}
}

func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *Set) Remove(item int) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *Set) Len() int {
	return len(s.List())
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
}

func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := []int{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func main() {
	//init
	s := New()

	s.Add(1)
	s.Add(1)
	s.Add(2)
	s.Clear()

	if s.IsEmpty() {
		fmt.Println("0 item")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.Has(2) {
		fmt.Println("2 does exist")
	} else {
		fmt.Println("2 does not exist")
	}

	s.Remove(2)
	s.Remove(3)
	fmt.Println("list all items: ", s.List())

	/*
		var t ISet
		t.Add(1)
		t.Add(2)
		t.Add(3)
		t.Has(1)
		t.Len()
		t.IsEmpty()
		t.list()
		t.Remove(1)
		fmt.Println("list all items: ", s.List())
		t.Clear()
		fmt.Println("list all items: ", s.List())
	*/

}
