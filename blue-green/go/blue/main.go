package main

import (
	"html/template"
	"log"
	"net/http"
)

type HomeViewModel struct {
	Title string
}

func main() {
	log.Println("Start ...")
	// setup
	templates := populateTemplates()
	start(templates)
	// start server
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Println("Error ...")
		log.Fatal(err)
	}
	log.Println(":9001")
}

func start(templates map[string]*template.Template) {
	template := templates["index"]
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var model = &HomeViewModel{Title: "Blue App"}
		template.Execute(w, model)
	})

	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/index.html"))
	tmpl := template.Must(layout.Clone())
	result["index"] = tmpl
	return result
}
