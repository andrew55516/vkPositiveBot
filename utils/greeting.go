package utils

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"unicode"
)

var searchStrings = map[string]string{
	"birthday": "https://pozdravok.com/pozdravleniya/den-rozhdeniya/",
	"new year": "https://pozdravok.com/pozdravleniya/prazdniki/noviy-god/",
}

func GetRandomGreeting(searchStr string) (string, error) {
	r, err := http.Get(searchStrings[searchStr])
	defer r.Body.Close()
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return "", err
	}

	greetingNodes := doc.Find(`p[class="sfst"]`).Nodes

	greeting := greetingNodes[rand.Intn(len(greetingNodes))]
	greetingStr := doc.Find(fmt.Sprintf(`#%s`, greeting.Attr[0].Val)).Text()

	decoder := charmap.Windows1251.NewDecoder()
	reader := transform.NewReader(bytes.NewReader([]byte(greetingStr)), decoder)
	decoded, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	greetingStr = string(decoded)

	builder := strings.Builder{}

	for _, c := range greetingStr {
		if unicode.IsUpper(c) {
			builder.WriteString("\n")
		}
		builder.WriteString(string(c))
	}

	greetingStr = builder.String()

	return greetingStr, nil
}
