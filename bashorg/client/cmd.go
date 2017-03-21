package main

import (
	"flag"
	"fmt"
	"sort"

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
		sort.Slice(quotes, func(i, j int) bool { return quotes[i].Votes < quotes[j].Votes })
	}
	for _, quote := range quotes {
		fmt.Printf("%s\n", quote)
	}
}
