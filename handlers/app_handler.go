package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/rahul-golang/crawler/services"
)

//AppHandlers handlers
type AppHandlers struct {
	service services.AppService
}

//NewAppHandlers creates dependancy
func NewAppHandlers(service services.AppService) *AppHandlers {
	return &AppHandlers{service: service}
}

//ServeApplication servres html page handlers
func (appHandlers *AppHandlers) ServeApplication(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/layout.html"))
	tmpl.Execute(w, nil)
}

//GetURLCount handler for crowl application
func (appHandlers *AppHandlers) GetURLCount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tmpl := template.Must(template.ParseFiles("templates/layout.html"))
	uri := r.FormValue("url_text")
	resp, err := http.Get(uri)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	charcount, err := appHandlers.service.CalWordCount(ctx, body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(charcount)
	tmpl.Execute(w, charcount)
}
