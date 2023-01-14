package main

import (
	"fmt"
	"./morestrings"
	"github.com/dougwatson/bar"
)
func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, world."))
	println("print from remote imported bar package=",bar.Hello())
}
