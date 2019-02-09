package main

import (
	"fmt"
	"time"
)

type SomeStruct struct {
	values map[string]int
}

func (s *SomeStruct) Count(key string) {
	value, ok := s.values[key]
	if !ok {
		s.values[key] = 1
	} else {
		s.values[key] = value + 1
	}
}

func main() {
	s := &SomeStruct{
		values: make(map[string]int),
	}

	for i := 0; i < 8; i++ {
		go s.Count("foo")
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("%d\n", s.values["foo"])
}
