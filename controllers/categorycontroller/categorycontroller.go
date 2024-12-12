package categorycontroller

import (
	"fp-pbkk2/models/categorymodel"
	"fp-pbkk2/entities"
	"net/http"
	"text/template"
	"time"
	"strconv"
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
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/edit.html")
		if err != nil {
			panic(err)
		}
		
		//konversi id string menjadi int 
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		category := categorymodel.Detail(id)
		data := map[string]any{
			"category": category,
		}

		temp.Execute(w, data)
	}


	if r.Method == "POST" {
		var category entities.Category
	
		// Konversi string id menjadi int
		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
	
		category.Name = r.FormValue("name")
		category.UpdatedAt = time.Now()
	
		// Memanggil fungsi Update di model
		if ok := categorymodel.Update(id, category); !ok {
			// Jika gagal update, redirect kembali ke halaman sebelumnya
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}
	
		// Jika data berhasil di-update
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
	
}

func Delete(w http.ResponseWriter, r *http.Request) {
   //konversi id string menjadi int 
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := categorymodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)

}
