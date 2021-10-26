package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/wayneashleyberry/rot13/pkg/rot13"
)

func main() {
	var decode bool
	flag.BoolVar(&decode, "decode", false, "...")
	flag.Parse()

	if decode {
		fmt.Println(rot13.Decode(strings.Join(flag.Args(), " ")))
		return
	}

	fmt.Println(rot13.Encode(strings.Join(flag.Args(), " ")))
}
