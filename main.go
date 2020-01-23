package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	//r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/temp", Static)
	r.HandleFunc("/test", Gettest)
	fmt.Println("starting server at 8 :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

//Gettest func..
func Gettest(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	setupResponse(&w, r)
	fmt.Fprintf(w, "first page2\n")
}

//Static func..
func Static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/dist/index.html")
	//http.StripPrefix("/", ))
}
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
