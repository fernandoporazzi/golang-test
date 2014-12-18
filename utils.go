package main

import (
  "net/http"
)

func setStaticDir() {
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))
}
