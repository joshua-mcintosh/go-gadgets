package bashorg

import (
"fmt"
"net/http"

"github.com/go-xmlpath/xmlpath"
)

const(
	bashRandomOne = "http://bash.org/?random1"
	quoteSelector = "/html/body/center[1]/table/tbody/tr/td[1]/p[@class=\"qt\"]"
	mdSelector = "preceding::p[@class=\"quote\"]"
)

var(
	quotePath = xmlpath.MustCompile(quoteSelector)
	mdPath = xmlpath.MustCompile(mdSelector)
	idPath = xmlpath.MustCompile("a[@title=\"Permanent link to this quote.\"]")
	votePath = xmlpath.MustCompile("font[1]")
	
)

type bashOrg struct{
	http.Client
}

type quote struct{
	id string
	votes string
	text string
}

func (q quote) String() string {
	return fmt.Sprintf("Id: %s -- Votes: %s\n----------\n%s\n", q.id, q.votes, q.text)
}

func NewBashOrg() bashOrg {
	return bashOrg{}
}

func (b *bashOrg) GetRandom() ([]quote, error) {
	qRet := make([]quote, 0)
	resp, err := b.Get(bashRandomOne)
	if err != nil {
		return qRet, err
	}
	defer resp.Body.Close()

	root, err := xmlpath.ParseHTML(resp.Body)
	if err != nil {
		return qRet, err
	}

	quotes := quotePath.Iter(root)
	for quotes.Next() {
		qt := quotes.Node()
		md := mdPath.Iter(qt)
		if !md.Next() {
			break
		}

		id, ok := idPath.String(md.Node())
		if !ok {
			break
		}

		votes, ok := votePath.String(md.Node())
		if !ok {
			break
		}
		q:= quote{id: id, votes:votes, text:qt.String()}
		qRet = append(qRet, q)
	}
	return qRet, nil
}
