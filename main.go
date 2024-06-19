package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oshaposhnyk/lenslocked/controllers"
	"github.com/oshaposhnyk/lenslocked/templates"
	"github.com/oshaposhnyk/lenslocked/views"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"layout.gohtml", "home.gohtml",
	))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"layout.gohtml", "contact.gohtml",
	))))

	r.Get("/faq", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"layout.gohtml", "faq.gohtml",
	))))

	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"layout.gohtml", "singup.gohtml",
	))
	r.Get("/singup", usersC.New)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
