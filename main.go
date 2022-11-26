package main

import (
	"7-routing/handler"
	"fmt"
	"net/http"

	"github.com/gorilla/mux" //npm routing
)

func main() {
	route := mux.NewRouter() // buat router dengan mux
	// port := "5000"

	// route untuk menginisialisasi folder public agar dapat dibaca
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	route.HandleFunc("/", handler.HandleHome).Methods("GET") // routing jadikan function home
	route.HandleFunc("/contact", handler.HandleContact).Methods("GET")
	route.HandleFunc("/project", handler.HandleProject).Methods("GET")
	route.HandleFunc("/add-project", handler.HandleAddProject).Methods("POST")
	route.HandleFunc("/project-detail/{id}", handler.HandleDetailProject).Methods("GET")
	route.HandleFunc("/delete/{id}", handler.HandleDeleteProject).Methods("GET")

	fmt.Println("Server sedang berjalan di port 3000 ")
	http.ListenAndServe("Localhost:3000", route) // panggil untuk dapat diakses di browser par1: string, par2: route
}
