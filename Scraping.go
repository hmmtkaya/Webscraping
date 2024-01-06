package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main () {
	res,_ := http.Get("https://www.sinefil.com/otodidakt/all/watchlist/63",)
	if res.StatusCode != 200 {
		fmt.Println("error", res.StatusCode)
		return
	}
	doc,_ := goquery.NewDocumentFromReader(res.Body)

	doc.Find(".hero", ).Each(func(i int, selection *goquery.Selection) {
		title := selection.Find("a").Text()
		title = strings.TrimSpace(title)
		fmt.Println(i+1, title)
		
	})

}