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

// Create struct, struct berfungsi untuk membuat struktur dari tipe data
type Project struct {
	ProjectName string
	StartDate   string
	EndDate     string
	Desc        string
	Tech        []string
	Image       string
}

// Buat array of object sebagai local storage
var DataProjects = []Project{
	{
		ProjectName: "Dumbways 2022",
		StartDate:   "2022-11-24",
		EndDate:     "2022-12-24",
		Desc:        "Halo Dumbways",
		Tech:        []string{"node", "angular", "react", "typescript"},
		Image:       "public/assets/img/saitama.png",
	},
}

func HandleHome(w http.ResponseWriter, r *http.Request) { //ResponseWriter: untuk menampilkan data, Request: untuk menambahkan data
	w.Header().Set("Content-type", "text/html; charset=utf-8") // Header berfungsi untuk menampilkan data. Data yang ditamplikan "text-html" /"json" / dll

	tmpt, err := template.ParseFiles("views/index.html") // template.ParseFiles berfungsi memparsing file yang disisipkan sebagai parameter

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	// Buat map, kemudian panggil local storage dengan id tertentu dan simpan local storage sebagai value dari "DataProjects"
	DataProjects := model.DataProjects
	dataProject := map[string]interface{}{
		"DataProjects": DataProjects,
	}
	// Kemudian tampilkan seluruh isi dari dari local storage
	tmpt.Execute(w, dataProject) // Execute berfungsi untuk mengeksekusi / menampilkan data dan harus ada 2 parameter (respon, Data)
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

	// Panggil struct kemudian tampung kedalam variabel dan masukkan data yang sudah ditangkap dengan PostForm.Get() ke dalam object
	newData := model.Project{
		ProjectName: projectName,
		StartDate:   startDate,
		EndDate:     endDate,
		Desc:        desc,
		Tech:        checkboxs,
	}

	// Kemudian panggil local storage dan append object newData
	model.DataProjects = append(model.DataProjects, newData)

	// Panggil method redirect agar Setelah data dikirim, maka routing akan berpindah ke halaman index
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	fmt.Println("Data telah berhasil ditambahkan")
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
	DataProjects := model.DataProjects
	dataProject := map[string]interface{}{
		"DataProjects": DataProjects[id],
	}

	tmpt.Execute(w, dataProject)
}

func HandleDeleteProject(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	model.DataProjects = append(model.DataProjects[:id], model.DataProjects[id+1:]...)

	http.Redirect(w, r, "/", http.StatusFound)
	fmt.Println("Data dengan ID ke", id, " berhasil dihapus")
}
