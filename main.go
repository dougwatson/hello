package main

import (
	"fmt"
	"os"
	"github.com/dougwatson/hello/morestrings"
	"github.com/dougwatson/bar"
)
func main() {
	word:="Hello,World"
	if len(os.Args)>1{
		word=os.Args[1]
	}
	fmt.Println(morestrings.ReverseRunes(word))
	println("import pkg:",bar.Hello())
}
