package main

import (
	"fp-pbkk2/config"
	"fp-pbkk2/controllers/categorycontroller"
	"fp-pbkk2/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	defer config.DB.Close()

	//Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	//Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server berjalan di port 8080")
	http.ListenAndServe(":8080", nil)
}
