package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"log"
	"net/http"
)

type footBallBlog struct {
	Author   string
	Contents string
	Topics   string
}

var Database []footBallBlog

//type Datastructure struct {
//	Datastructure []Sportsblog
//}

//var Data = Datastructure{Datastructure: []Sportsblog{}}
//var templates *template.Template

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", HandleHomepage)
	r.Post("/postblog", HandlePost)
	r.Delete("/delete", HandleDeletePost)

	fmt.Println("Server Launched") //notice to show the server has been launched at the port below
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}

}

//display page
func HandleHomepage(w http.ResponseWriter, r *http.Request) {

	templ := template.Must(template.ParseFiles("Template/home.html"))
	err := templ.Execute(w, Database)
	if err != nil {
		log.Fatal(err)
	}

}

//handle the post
func HandlePost(w http.ResponseWriter, r *http.Request) {

	//GET THE INFORMATION FROM THE HTML FORM
	GetAuthor := r.FormValue("username")
	GetTopic := r.FormValue("topic")
	GetContent := r.FormValue("content")

	//INSTANTIATE OR POPULATE OR ASSIGN A VALUE TO DATA
	data := footBallBlog{
		Author:   GetAuthor,
		Contents: GetContent,
		Topics:   GetTopic,
	}
	Database = append(Database, data) // append the data to the database. This will run the add function.
	fmt.Println(Database)

	//PUSH THE DATA INTO THE TEMPLATE
	templ := template.Must(template.ParseFiles("Template/home.html"))
	err := templ.Execute(w, Database)
	if err != nil {
		log.Fatal(err)
	}
}

//Delete post
func HandleDeletePost(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("Template/delete.html"))
	err := templ.Execute(w, Database)
	if err != nil {
		log.Fatal(err)
	}
}
