package handler

import (
	"7-routing/model"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func HandleHome(w http.ResponseWriter, r *http.Request) { //ResponseWriter: untuk menampilkan data, Request: untuk menambahkan data
	w.Header().Set("Content-type", "text/html; charset-utf-8") // Header berfungsi untuk menampilkan data. Data yang ditamplikan "text-html" /"json" / dll

	tmpl, err := template.ParseFiles("views/index.html") // template.ParseFiles berfungsi memparsing file yang disisipkan sebagai parameter

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	fmt.Println((model.DataProjects))
	dataProject := map[string]interface{}{
		"DataProjects": model.DataProjects,
	}

	tmpl.Execute(w, dataProject) // Execute berfungsi untuk mengeksekusi / menampilkan data dan harus ada 2 parameter (respon, Data)

}

func HandleContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset-utf-8")

	result, err := template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	} else {
		result.Execute(w, nil)
	}
}

func HandleProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset-utf-8")
	tmpt, err := template.ParseFiles("views/project.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	} else {
		tmpt.Execute(w, nil)
	}
}

func HandleDetailProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset-utf-8")
	tmpt, err := template.ParseFiles("views/project-detail.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	// Tangkap id dari blog
	id, _ := strconv.Atoi(mux.Vars(r)["id"]) // strconv.Atoi untuk konversi string ke int.  mux.Vars() berfungsi untuk menangkap id dan mengembalikan 2 nilai parameter result dan error

	// Buat variable untuk menampung data struct
	var dataProjectDetail = model.Project{}

	for i, data := range model.DataProjects {
		if i == id {
			dataProjectDetail = model.Project{
				ProjectName: data.ProjectName,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Desc:        data.Desc,
			}
		}
	}

	// Buat slice untuk menampung variabel dataProjectDetail
	dataProject := map[string]interface{}{
		"Data": dataProjectDetail,
	}

	tmpt.Execute(w, dataProject)
}

func HandleAddProject(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r) // r berisi seluruh data form

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	// Tangkap data form dengan method PostForm.Get
	projectName := r.PostForm.Get("projectName")
	startDate := r.PostForm.Get("startDate")
	endDate := r.PostForm.Get("endDate")
	desc := r.PostForm.Get("desc")

	//  Buat variabel untuk menampung data checkbox
	var checkboxs []string

	// Jika didalam form checkboxs ada value-nya, maka append ke array checkboxs
	if r.FormValue("node") != "" {
		checkboxs = append(checkboxs, r.FormValue("node"))
	}
	if r.FormValue("angular") != "" {
		checkboxs = append(checkboxs, r.FormValue("angular"))
	}
	if r.FormValue("react") != "" {
		checkboxs = append(checkboxs, r.FormValue("react"))
	}
	if r.FormValue("typescript") != "" {
		checkboxs = append(checkboxs, r.FormValue("typescript"))
	}

	// Buat object dengan menginisialisasi key berdasarkan struct
	newData := model.Project{
		ProjectName: projectName,
		StartDate:   startDate,
		EndDate:     endDate,
		Desc:        desc,
		Tech:        checkboxs,
	}

	// Buat penampung datas kemudian append / masukkan object newData ke dalam slice DataProject
	model.DataProjects = append(model.DataProjects, newData)

	// Panggil method redirect agar Setelah data dikirim, maka routing akan berpindah ke halaman index
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	index, _ := strconv.Atoi(mux.Vars(r)["index"])

	model.DataProjects = append(model.DataProjects[:index], model.DataProjects[index+1:]...)

	http.Redirect(w, r, "/", http.StatusAccepted)
}
