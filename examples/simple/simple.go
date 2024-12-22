package main

import (
	"log"

	"github.com/skazanyNaGlany/go-mustget"
)

func main() {
	someMap := map[string]any{
		"key": "some value",
	}

	value := mustget.MustGet(someMap, "key").(string)

	log.Println(value)
}
