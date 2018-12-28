package basic_date

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"testing"
)

func TestQuery(t *testing.T) {
	movieList := make([]map[string]string, 0)

	resp, err := http.Get("https://www.vmovier.com/")
	if err != nil {
		panic(err)
	}
	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("#post-list li").Each(func(i int, selection *goquery.Selection) {
		title, _ := selection.Find("li .index-img").Attr("title")
		imgpath, _ := selection.Find("li img").Attr("src")
		username := selection.Find("li .user").Text()
		movieList = append(movieList, map[string]string{"title": title, "path": imgpath, "name": username})
	})

	jsonhtml, err := json.Marshal(movieList)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonhtml))
}
