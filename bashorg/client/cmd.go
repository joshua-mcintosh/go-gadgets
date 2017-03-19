package main

import (
	"fmt"
	"sort"

	bo "github.com/joshua-mcintosh/go-gadgets/bashorg"
)

func main() {
	b := bo.NewBashOrg()
	quotes, err := b.GetRandom()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	sort.Sort(bo.QuotesByVote(quotes))
	for _, quote := range quotes {
		fmt.Printf("%s\n", quote)
	}
}
