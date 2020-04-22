package main

import (
  	"fmt"
	"github.com/ede-n/golang/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
    fmt.Println(morestrings.ReverseRunes("Hello, World\n"))
    fmt.Println(cmp.Diff("Hello World", "Hello Go"))

}
