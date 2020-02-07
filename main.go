package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})
	http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {

		uri := r.FormValue("url_text")
		resp, err := http.Get(uri)
		if err != nil {
			return
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		//fmt.Println(string(body))
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
		// fmt.Print(array)
		charcount := map[string]int{}
		for _, c := range array {
			charcount[string(c)]++
		}
		// fmt.Println(charcount)
		tmpl.Execute(w, charcount)
	})
	fmt.Println(http.ListenAndServe(":8000", nil))
}
