package main

import (
	"fmt"
	"strings"
)

type Set map[int]struct{}

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(value int) {
	s[value] = struct{}{}
}

func (s Set) Remove(value int) {
	delete(s, value)
}

func (s Set) Contains(value int) bool {
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
		elements = append(elements, fmt.Sprintf("%d", key))
	}
	return "{" + strings.Join(elements, ", ") + "}"
}

func main() {
	set1 := NewSet()
	set1.Add(1)
	set1.Add(24)
	set1.Add(13)
	set1.Add(5)
	set1.Add(0)
	set1.Add(124)

	set2 := NewSet()
	set2.Add(3)
	set2.Add(4)
	set2.Add(124)
	set2.Add(1)

	intersection := set1.Intersection(set2)
	fmt.Println("Пересечение множеств:", intersection.Format())

}
