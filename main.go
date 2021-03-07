package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//IndexHandler fubction
func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}
	return http.HandlerFunc(fn)
}
func newRouter() *mux.Router {
	r := mux.NewRouter()

	//	staticFiles := http.FileServer(http.Dir("frontend/dist/static"))

	//	r.PathPrefix("/css").Handler(staticFiles)
	//	r.PathPrefix("/js").Handler(staticFiles)
	r.PathPrefix("/").HandlerFunc(IndexHandler("frontend/dist/index.html"))
	r.HandleFunc("/test", Gettest)
	return r
}

func main() {
	r := mux.NewRouter()
	// It's important that this is before your catch-all route ("/")
	api := r.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/users", GetUsersHandler).Methods("GET")
	// Serve static assets directly.
	r.PathPrefix("/static").Handler(http.FileServer(http.Dir("frontend/dist")))
	// Catch-all: Serve our JavaScript application's entry-point (index.html).
	r.PathPrefix("/").HandlerFunc(IndexHandler("frontend/dist/index.html"))

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, r),
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}

//Gettest func..
func Gettest(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	fmt.Println("Gettest Request")
	//setupResponse(&w, r)
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

//GetUsersHandler function
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"id":        "123",
		"timeStamp": time.Now().Format(time.RFC3339),
	}
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write(b)
}
