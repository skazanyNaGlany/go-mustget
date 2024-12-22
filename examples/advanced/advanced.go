package main

import (
	"log"

	"github.com/skazanyNaGlany/go-mustget"
)

type SomeExampleStruct struct{}

func (s *SomeExampleStruct) SomeMethod() {
	log.Println("SomeMethod called")
}

func main() {
	someMap := map[**SomeExampleStruct]any{}

	key1 := &SomeExampleStruct{}
	key2 := &SomeExampleStruct{}
	key3 := &SomeExampleStruct{}

	someMap[&key1] = "some value 1"
	someMap[&key2] = "some value 2"
	someMap[&key3] = "some value 3"

	value := mustget.MustGet(someMap, &key2).(string)

	log.Println(value)
}
