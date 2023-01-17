package main

import (
	"fmt"
	"github.com/dougwatson/hello/morestrings"
	"github.com/dougwatson/bar"
)
func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, world."))
	println("import pkg:",bar.Hello())
}
