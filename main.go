package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type ContactDetails struct {
	Job string
	Day string
	Who string
}

func main() {
	http.HandleFunc("/", Home)
	fmt.Println("Starting server on port 8080")

	// Start the sever
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("form.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := ContactDetails{
		Job: r.FormValue("job"),
		Day: r.FormValue("day"),
		Who: r.FormValue("who"),
	}

	_ = details

	tmpl.Execute(w, struct{ Success bool }{true})

}
