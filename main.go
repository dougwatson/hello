package main

import (
	"fmt"
	. "fs/fs" // This import is required to use the easyFS filesystem
	"os"

	"github.com/dougwatson/xbar"
	"github.com/dougwatson/xhello/morestrings"
	"github.com/mattn/go-runewidth"
)

var fs = easyFS

func main() {
	x := runewidth.StringWidth("つのだ☆HIRO")
	fmt.Printf("x=%v\n", x)
	word := "Hello,World"
	if len(os.Args) > 1 {
		word = os.Args[1]
	}
	fmt.Println(morestrings.ReverseRunes(word))
	println("import pkg:", xbar.Hello())
	/************************************/

	b, err := fs.ReadFile("home/src/github.com/dougwatson/xhello/go.mod")
	if err != nil {
		println("error: " + err.Error())
	}
	println("b=", string(b))
}
