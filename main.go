package main

import (
	"fmt"
	"os"
	"github.com/dougwatson/xhello/morestrings"
	"github.com/dougwatson/xbar"
	"github.com/gocoderpro/easyfs"
)
func main() {
	word:="Hello,World"
	if len(os.Args)>1{
		word=os.Args[1]
	}
	fmt.Println(morestrings.ReverseRunes(word))
	println("import pkg:",xbar.Hello())
	fs:=GetFS()
	b,err:=fs.ReadFile("home/src/github.com/dougwatson/xhello")
	if err!=nil {
	   println("error: "+err)
	}
	println("b=",string(b))
	
}
