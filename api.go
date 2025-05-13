package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func getfile(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
	}
	defer r.Body.Close()

	err = os.WriteFile("./upload.txt", data, 0644)
	if err != nil {
		http.Error(w, "failed to save file", http.StatusInternalServerError)
		return
	}
	fmt.Println("File recieved successfully!")
	w.Write([]byte("File recieved successfully!"))
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Success"))
}
func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", getfile)
	r.Get("/", test)
	return r
}
func (app *application) run(r *chi.Mux) error {
	srv := http.Server{
		Addr:    app.config.addr,
		Handler: r,
	}
	log.Print("Server running on port 8080")
	return srv.ListenAndServe()
}
