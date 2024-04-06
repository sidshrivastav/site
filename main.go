package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
  "path/filepath"
)

// Template directory
var templatesDir = "templates"

// Data structure for home
type HomeData struct {
    LuckyNumber int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Parse the home template file
		tmpl, err := template.ParseFiles(filepath.Join(templatesDir, "layout.html"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

    // Data to show
		data := HomeData{
      LuckyNumber: 1,
		}

		// Execute the home template, passing in the home data
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

  // Serve Static File
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
