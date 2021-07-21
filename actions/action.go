package actions

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/pArtour/telegram-parser-bot/telegram"
	"log"
	"net/http"
)

func Action(tag string) {
	url := fmt.Sprintf("https://habr.com/ru/hub/%s", tag)
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatalf(err.Error())
	}

	title := doc.Find(".tm-articles-list__item").Find(".tm-article-snippet__title_h2")
	link, _ := title.Find("a").Attr("href")
	lintText, _ := title.Find("a").Find("span").Html()
	fmt.Println(link, lintText)
	habrUrl := "https://habr.com"
	text := fmt.Sprintf(`<b>Habr - %s</b>: <a href\=\"%s%s\">%s</a>`, tag, habrUrl, link, lintText)

	fmt.Println(text)

	telegram.SendMessage(text)

}