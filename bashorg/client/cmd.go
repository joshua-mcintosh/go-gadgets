package main

import (
	"fmt"
	"github.com/joshua-mcintosh/go-gadgets/bashorg"
)

func main() {
	bo := bashorg.NewBashOrg()
	resp, err := bo.GetRandom()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, quote := range resp {
		fmt.Printf("%s\n", quote)
	}
}
