package main

import (
  "html/template"
  "net/http"
  "time"
)

func main() {
  setStaticDir()

  http.HandleFunc("/", index)
  http.HandleFunc("/about", about)
  http.HandleFunc("/service", service)
  http.ListenAndServe(":8080", nil)
}

func setStaticDir() {
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))
}

// Available templates.
var templates = template.Must(template.ParseFiles(
  "templates/index.html",
  "templates/about.html",
  "templates/service.html",
))

// data needed to render the template
type About struct {
  Description string
}

type Service  struct {
  Services string
}

type Page struct {
  About       *About
  Service     *Service
  Title       string
  CurrentYear string
}

// Render templates.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
  err := templates.ExecuteTemplate(w, tmpl+".html", p)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

// Handler to show the index page
func index(w http.ResponseWriter, r *http.Request) {
  year := time.Now().Local()
  page := Page{
    Title:       "Home",
    CurrentYear: year.Format("2006"),
  }
  renderTemplate(w, "index", &page)
}

// Handler to show about page
func about(w http.ResponseWriter, r *http.Request) {
  about := Page{
    Title: "About",
    About: &About{
        Description: "lorem ipsum",
      },
  }
  renderTemplate(w, "about", &about)
}

// Handler to show about page
func service(w http.ResponseWriter, r *http.Request) {
  service := Page{
    Title: "About",
    Service: &Service{
        Services: "this should be a list where I name the things I do",
      },
  }
  renderTemplate(w, "service", &service)
}