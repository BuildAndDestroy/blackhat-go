package bing

import (
	"archive/zip"
	"blackhat-go/scraper/metadata"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func Handler(i int, s *goquery.Selection) {
	url, ok := s.Find("a").Attr("href")
	if !ok {
		return
	}

	fmt.Printf("%d: %s\n", i, url)
	res, err := http.Get(url)
	if err != nil {
		return
	}
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()

	r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return
	}

	cp, ap, err := metadata.NewProperties(r)
	if err != nil {
		return
	}

	log.Printf(
		"%25s %25s - %s %s\n",
		cp.Creator,
		cp.LastModifiedBy,
		ap.Application,
		ap.GetMajorVersion())
}

func BingSearch() {
	if len(os.Args) != 3 {
		log.Fatalln("Missing required argument. Usage: main.go domain ext")
	}
	domain := os.Args[1]
	filetype := os.Args[2]

	q := fmt.Sprintf(
		"site:%s && filetype:%s && instreamset:(url title):%s",
		domain,
		filetype,
		filetype)
	search := fmt.Sprintf("http://www.bing.com/search?q=%s", url.QueryEscape(q))
	// doc, err := goquery.NewDocument(search)
	// if err != nil {
	// 	log.Panicln(err)
	// }
	resp, err := http.Get(search)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Panicln(err)
	}

	s := "html body div#b_content ol#b_results li.b_algo div.b_title h2"
	doc.Find(s).Each(Handler)
}
