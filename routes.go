package main

import (
  "html/template"
  "net/http"
  "time"
)

// Available templates.
var templates = template.Must(template.ParseFiles(
  "templates/index.html",
  "templates/about.html",
  "templates/service.html",
))

// Handler to show the index page
func index(w http.ResponseWriter, r *http.Request) {
  year := time.Now().Local()
  data := struct {
    Title string
    CurrentYear string
  }{
    "Home",
    year.Format("2006"),
  }
  templates.ExecuteTemplate(w, "index.html", &data)
}

// Handler to show about page
func about(w http.ResponseWriter, r *http.Request) {
  data := struct {
      Title string
      About string
    }{
      "About",
      "Lorem ipsum",
    }
  templates.ExecuteTemplate(w, "about.html", &data)
}

// Handler to show service page
func service(w http.ResponseWriter, r *http.Request) {
  data := struct {
      Title string
      Service string
    }{
      "Service",
      "this should be a list where I name the things I do",
    }
  templates.ExecuteTemplate(w, "service.html", &data)
}
