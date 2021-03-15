package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//go:embed assets
//go:embed templates
//go:embed content
var files embed.FS

func baseURL(w http.ResponseWriter, r *http.Request) {
	url := "/html/index.html"
	http.Redirect(w, r, url, http.StatusSeeOther)
}
func html(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page := strings.Replace(vars["page"], ".html", "", -1)
	tmpl := "templates/template.html"
	content := "content/" + page + ".html"
	title := "content/" + page + ".title"
	parsedTemplate, err := template.ParseFS(files, tmpl, content, title)
	if err != nil {
		log.Println("Error loading template for "+page+": ", err)
		return
	}
	err = parsedTemplate.Execute(w, page)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}
func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	fs := http.FileServer(http.FS(files))
	myRouter.HandleFunc("/", baseURL)
	myRouter.PathPrefix("/assets/").Handler(http.StripPrefix("/", fs))
	myRouter.HandleFunc("/html/{page}", html)
	log.Fatal(http.ListenAndServe(":8080", handlers.CombinedLoggingHandler(os.Stdout, myRouter)))
}
