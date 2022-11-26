package handler

import (
	"7-routing/connection"
	"7-routing/model"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func HandleHome(w http.ResponseWriter, r *http.Request) { //ResponseWriter: untuk menampilkan data, Request: untuk menambahkan data
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // Header berfungsi untuk menampilkan data. Data yang ditamplikan "text-html" /"json" / dll

	tmpt, err := template.ParseFiles("views/index.html") // template.ParseFiles berfungsi memparsing file yang disisipkan sebagai parameter

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	// Query data dari database
	dataProject, err := connection.Conn.Query(context.Background(), "SELECT id, project_name, start_date, end_date, description, technologies, image FROM tb_projects") // Query mengembalikan 2 nilai
	if err != nil {
		w.Write([]byte("Message :" + err.Error()))
		return
	}

	// // Panggil struct untuk menampung data dari database
	var result []model.Project

	// // Sebelum data ditampilkan looping terlebih dahulu
	for dataProject.Next() {
		var each = model.Project{}

		err := dataProject.Scan(&each.ID, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Desc, &each.Tech, &each.Image)
		if err != nil {
			fmt.Println("Message : ", err.Error())
			return
		}
		result = append(result, each)
	}

	// Buat map sebagai penampung data result
	data := map[string]interface{}{
		"DataProjects": result,
	}
	fmt.Println("data :", result)

	// Kemudian tampilkan seluruh data dari database
	w.WriteHeader(http.StatusOK)
	tmpt.Execute(w, data)
}

func HandleProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/project.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	} else {
		tmpt.Execute(w, nil)
	}
}

func HandleContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	result, err := template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	} else {
		result.Execute(w, nil)
	}
}

func HandleDetailProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/project-detail.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	// Tangkap id dari blog
	id, _ := strconv.Atoi(mux.Vars(r)["id"]) // strconv.Atoi untuk konversi string ke int.  mux.Vars() berfungsi untuk menangkap id dan mengembalikan 2 nilai parameter result dan error

	// Buat map, kemudian panggil local storage dengan id tertentu dan simpan local storage sebagai value dari "DataProjects"
	// Query data dari database
	dataProject, err := connection.Conn.Query(context.Background(), "SELECT id, project_name, start_date, end_date, description, technologies, image FROM tb_projects") // Query mengembalikan 2 nilai
	if err != nil {
		w.Write([]byte("Message :" + err.Error()))
		return
	}

	// // Panggil struct untuk menampung data dari database
	var result []model.Project

	// // Sebelum data ditampilkan looping terlebih dahulu
	for dataProject.Next() {
		var each = model.Project{}

		err := dataProject.Scan(&each.ID, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Desc, &each.Tech, &each.Image)
		if err != nil {
			fmt.Println("Message : ", err.Error())
			return
		}
		// Kemudian setiap data dimasukkan ke dalam result
		result = append(result, each)
	}

	// Buat map sebagai penampung data result
	data := map[string]interface{}{
		"DataProjects": result[id],
	}

	w.WriteHeader(http.StatusOK)
	tmpt.Execute(w, data)
}
