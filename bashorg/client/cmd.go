package main

import (
	"flag"
	"fmt"

	bo "github.com/joshua-mcintosh/go-gadgets/bashorg"
)

var(
	sortByVotes = flag.Bool("by-vote", false, "Sort the quotes by vote")
)

func main() {
	flag.Parse()

	b := bo.NewBashOrg()
	quotes, err := b.GetRandom()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	if *sortByVotes {
		quotes.SortByVote()
	}
	for _, quote := range quotes {
		fmt.Printf("%s\n", quote)
	}
}
