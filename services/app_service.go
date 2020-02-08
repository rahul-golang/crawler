package services

import (
	"context"
	"strings"

	"golang.org/x/net/html"
)

//AppServiceImpl application services
type AppServiceImpl struct {
}

//AppService **
type AppService interface {
	CalWordCount(context.Context, []byte) (map[string]int, error)
}

//NewAppServiceImpl **
func NewAppServiceImpl() AppService {

	return AppServiceImpl{}
}

//CalWordCount Calculate worrd count and return value
func (appService AppServiceImpl) CalWordCount(ctx context.Context, body []byte) (map[string]int, error) {

	final := ""

	domDocTest := html.NewTokenizer(strings.NewReader(string(body)))
	previousStartTokenTest := domDocTest.Token()
loopDomTest:
	for {
		tt := domDocTest.Next()
		switch {
		case tt == html.ErrorToken:
			break loopDomTest // End of the document,  done
		case tt == html.StartTagToken:
			previousStartTokenTest = domDocTest.Token()
		case tt == html.TextToken:
			if previousStartTokenTest.Data == "script" {
				continue
			}
			TxtContent := strings.TrimSpace(html.UnescapeString(string(domDocTest.Text())))
			if len(TxtContent) > 0 {
				final = final + TxtContent

			}
		}
	}

	array := strings.Split(final, " ")
	charcount := map[string]int{}
	for _, c := range array {
		charcount[string(c)]++
	}
	return charcount, nil
}
