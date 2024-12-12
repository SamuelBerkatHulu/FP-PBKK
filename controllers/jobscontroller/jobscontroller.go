package jobscontroller

import (
	"fp-pbkk2/entities"
	"fp-pbkk2/models/jobsmodel"
	"fp-pbkk2/models/categorymodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

// Index handles the listing of jobs
func Index(w http.ResponseWriter, r *http.Request) {
	jobs := jobsmodel.GetAll()
	data := map[string]any{
		"jobs": jobs,
	}

	temp, err := template.ParseFiles("views/jobs/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

// Detail handles showing details of a specific job
func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid job ID", http.StatusBadRequest)
		return
	}

	job := jobsmodel.Detail(id)
	data := map[string]any{
		"job": job,
	}

	temp, err := template.ParseFiles("views/jobs/detail.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Add handles the creation of a new job
func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/jobs/create.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
		}

		temp.Execute(w, data)
		return
	}

	if r.Method == http.MethodPost {
		var job entities.Job

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}

		vacancies, err := strconv.Atoi(r.FormValue("vacancies"))
		if err != nil {
			http.Error(w, "Invalid vacancies", http.StatusBadRequest)
			return
		}

		salary, err := strconv.ParseFloat(r.FormValue("salary"), 64)
		if err != nil {
			http.Error(w, "Invalid salary", http.StatusBadRequest)
			return
		}

		job.Title = r.FormValue("title")
		job.Category.Id = uint(categoryId)
		job.Vacancies = uint(vacancies)
		job.Description = r.FormValue("description")
		job.Salary = salary
		job.CreatedAt = time.Now()
		job.UpdatedAt = time.Now()

		if !jobsmodel.Create(job) {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/jobs", http.StatusSeeOther)
	}
}

// Edit handles editing an existing job
func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/jobs/edit.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid job ID", http.StatusBadRequest)
			return
		}

		job := jobsmodel.Detail(id)
		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
			"job":        job,
		}

		err = temp.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	if r.Method == http.MethodPost {
		var job entities.Job

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid job ID", http.StatusBadRequest)
			return
		}

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}

		vacancies, err := strconv.Atoi(r.FormValue("vacancies"))
		if err != nil {
			http.Error(w, "Invalid vacancies", http.StatusBadRequest)
			return
		}

		salary, err := strconv.ParseFloat(r.FormValue("salary"), 64)
		if err != nil {
			http.Error(w, "Invalid salary", http.StatusBadRequest)
			return
		}

		job.Title = r.FormValue("title")
		job.Category.Id = uint(categoryId)
		job.Vacancies = uint(vacancies)
		job.Description = r.FormValue("description")
		job.Salary = salary
		job.UpdatedAt = time.Now()

		if !jobsmodel.Update(id, job) {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/jobs", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idString)
    if err != nil {
        panic(err)
    }

    // Panggil model untuk menghapus job berdasarkan ID
    if err := jobsmodel.Delete(id); err != nil {
       panic(err)
    }

    // Redirect ke halaman daftar pekerjaan
    http.Redirect(w, r, "/jobs", http.StatusSeeOther)
}