package main

import (
	"fmt"
	"os"

	"github.com/dougwatson/xbar"
	"github.com/dougwatson/xhello/morestrings"

	"github.com/gocoderpro/easyfs"
)

func main() {
	word := "Hello,World"
	if len(os.Args) > 1 {
		word = os.Args[1]
	}
	fmt.Println(morestrings.ReverseRunes(word))
	println("import pkg:", xbar.Hello())
	/************************************/

	fs := easyfs.NewFS()
	_, err := fs.ReadFile("home/src/github.com/dougwatson/xhello/go.mod")
	if err != nil {
		println("error: " + err.Error())
	}
	//println("b=", string(b))
}
