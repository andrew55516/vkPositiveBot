package utils

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func GetRandomPicture(searchStr string) (string, error) {
	searchStr = strings.Replace(searchStr, " ", "-", -1)

	r, err := http.Get("https://unsplash.com/s/photos/" + searchStr)
	defer r.Body.Close()
	if err != nil {
		log.Println(err)
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	imgNodes := doc.Find(`div[class="MorZF"]`).Nodes

	rand.Seed(time.Now().UnixNano())

	var imgURL string
	for !strings.HasPrefix(imgURL, "https://") {
		imgURL = imgNodes[rand.Intn(len(imgNodes))].FirstChild.Attr[3].Val

	}

	return imgURL, nil
}
