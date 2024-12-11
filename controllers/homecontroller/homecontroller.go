package homecontroller

import (
	"net/http"
	"text/template"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	// Memuat template
	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	// Merender template
	err = temp.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}
