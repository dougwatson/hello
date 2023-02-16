package main

import (
	"fmt"
	_ "fs/fs" // This import is required for the filesystem to be available
	"os"

	"github.com/dougwatson/xbar"
	"github.com/dougwatson/xhello/morestrings"
)

func main() {
	word := "Hello,World"
	if len(os.Args) > 1 {
		word = os.Args[1]
	}
	fmt.Println(morestrings.ReverseRunes(word))
	println("import pkg:", xbar.Hello())
	/************************************/

	filesystem := fs.val
	_, err := filesystem.ReadFile("home/src/github.com/dougwatson/xhello/go.mod")
	if err != nil {
		println("error: " + err.Error())
	}
	//println("b=", string(b))
}
