package utils

import (
	"github.com/PuerkitoBio/goquery"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func GetRandomQuote() (string, error) {
	r, err := http.Get("https://ru.citaty.net/motivatsionnye-tsitaty/")
	defer r.Body.Close()
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return "", err
	}

	quoteNodes := doc.Find(`p[class="blockquote-text"]`).Nodes

	rand.Seed(time.Now().UnixNano())
	quote := quoteNodes[rand.Intn(len(quoteNodes))].FirstChild.Attr[1].Val

	quote = strings.TrimPrefix(quote, "Подробная цитата ")

	return quote, nil

}
