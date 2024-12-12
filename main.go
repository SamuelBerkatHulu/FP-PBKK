package main

import (
	"fp-pbkk2/config"
	"fp-pbkk2/controllers/categorycontroller"
	"fp-pbkk2/controllers/homecontroller"
	"fp-pbkk2/controllers/jobscontroller"
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


	//job
	http.HandleFunc("/jobs", jobscontroller.Index)
	http.HandleFunc("/jobs/add", jobscontroller.Add)
	http.HandleFunc("/jobs/detail", jobscontroller.Detail)
	http.HandleFunc("/jobs/edit", jobscontroller.Edit)
	http.HandleFunc("/jobs/delete", jobscontroller.Delete)


	log.Println("Server berjalan di port 8080")
	http.ListenAndServe(":8080", nil)
}
