package bashorg

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/go-xmlpath/xmlpath"
)

const (
	bashRandomOne = "http://bash.org/?random1"
	quoteSelector = `/html/body/center[1]/table/tbody/tr/td[1]/p[@class="qt"]`
	mdSelector    = `preceding::p[@class="quote"]`
)

var (
	quotePath = xmlpath.MustCompile(quoteSelector)
	mdPath    = xmlpath.MustCompile(mdSelector)
	idPath    = xmlpath.MustCompile(`a[@title="Permanent link to this quote."]`)
	votePath  = xmlpath.MustCompile("font[1]")
)

type bashOrg struct {
	http.Client
}

type quote struct {
	id    int
	votes int
	text  string
}

type quotes []quote

func (q quotes) Len() int           { return len(q) }
func (q quotes) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q quotes) Less(i, j int) bool { return q[i].id < q[j].id }

type QuotesByVote quotes

func (q QuotesByVote) Len() int           { return len(q) }
func (q QuotesByVote) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q QuotesByVote) Less(i, j int) bool { return q[i].votes < q[j].votes }

func (q quote) String() string {
	return fmt.Sprintf("Id: %d -- Votes: %d\n----------\n%s\n", q.id, q.votes, q.text)
}

func NewBashOrg() bashOrg {
	return bashOrg{}
}

func (b *bashOrg) GetRandom() ([]quote, error) {
	qRet := make(quotes, 0)
	resp, err := b.Get(bashRandomOne)
	if err != nil {
		return qRet, err
	}
	defer resp.Body.Close()

	root, err := xmlpath.ParseHTML(resp.Body)
	if err != nil {
		return qRet, err
	}

	qs := quotePath.Iter(root)
	for qs.Next() {
		qt := qs.Node()
		md := mdPath.Iter(qt)
		if !md.Next() {
			break
		}

		idStr, ok := idPath.String(md.Node())
		if !ok {
			break
		}
		id, err := strconv.Atoi(strings.Trim(idStr, "#"))
		if err != nil {
			return qRet, err
		}

		votesStr, ok := votePath.String(md.Node())
		if !ok {
			break
		}
		votes, err := strconv.Atoi(strings.Trim(votesStr, "#"))
		if err != nil {
			return qRet, err
		}

		q := quote{
			id:    id,
			votes: votes,
			text:  qt.String(),
		}
		qRet = append(qRet, q)
	}
	sort.Sort(qRet)
	return qRet, nil
}
