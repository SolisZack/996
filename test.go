package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a float32 = 3.0
	var b int = 3
	c := "sad"

	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), c)
}
