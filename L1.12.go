package main

import (
	"fmt"
	"strings"
)

type Set map[string]struct{}

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(value string) {
	s[value] = struct{}{}
}

func (s Set) Remove(value string) {
	delete(s, value)
}

func (s Set) Contains(value string) bool {
	_, exists := s[value]
	return exists
}

func (s Set) Intersection(other Set) Set {
	result := NewSet()
	for key := range s {
		if other.Contains(key) {
			result.Add(key)
		}
	}
	return result
}

func (s Set) Union(other Set) Set {
	result := NewSet()
	for key := range s {
		result.Add(key)
	}
	for key := range other {
		result.Add(key)
	}
	return result
}
func (s Set) Format() string {
	var elements []string
	for key := range s {
		elements = append(elements, fmt.Sprintf("%s", key))
	}
	return "{" + strings.Join(elements, ", ") + "}"
}
func main() {
	strSet := NewSet()
	st := "cat, cat, dog, cat, tree"
	starr := strings.Split(st, ", ")
	for _, i := range starr {
		strSet.Add(i)
	}
	fmt.Println(strSet.Format())
}
