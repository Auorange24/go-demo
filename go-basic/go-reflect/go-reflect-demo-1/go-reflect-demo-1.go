package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f float32 = 1.2
	fmt.Printf("%v\n", reflect.TypeOf(f))
}