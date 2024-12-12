package categorycontroller

import (
	"fp-pbkk2/models/categorymodel"
	"fp-pbkk2/entities"
	"net/http"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)

}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/create.html")
		if err != nil{
			panic(err)
		}
	temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var category entities.Category

		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		//Menankap inputas user dari categorymodel.go
		// function Create
		if ok := categorymodel.Create(category); !ok {
			//jika imput tidak berhail maka kembali ke tempat semula
			temp, _ := template.ParseFiles("views/category/create.html")
			temp.Execute(w, nil)
		}


		//Berhasil input maka
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
