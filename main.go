package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/base.html"))
	if r.Method == "GET" {
		tmpl.Execute(w, "my_data")
	} else {
		tmpl.Execute(w, "my_data")
	}
}

func me(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/likes.html"))
	vars := mux.Vars(r)
	id := vars["uuid"]
	g := getEngine()
	var count Count
	data := make(map[string]interface{})

	if r.Method == "GET" {
		// get amount of loves form DB

		g.Where("username == ?", id).Find(&count)
		data["count"] = count.Counts
		data["uuid"] = id
		tmpl.Execute(w, data)
	} else if r.Method == "POST" {
		data["uuid"] = id
		// allow me to love him. Make a form, or a button that triggers ajax call!

	}

}

func generateHugs(w http.ResponseWriter, r *http.Request) {
	id := uuid.New().String()
	data := struct {
		ID string `json:"id"`
	}{id}
	res, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func main() {
	router := mux.NewRouter()
	log.Printf("Serving on: http://localhost:8912")
	router.HandleFunc("/me", generateHugs)
	router.HandleFunc("/", index)
	router.HandleFunc("/{uuid}", me)

	http.Handle("/", router)

	http.ListenAndServe(":8912", nil)
}

type Count struct {
	gorm.Model
	Counts int
	User   string
}

func getEngine() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Count{})
	return db
}
