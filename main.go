package main

import (
  "net/http"
)

func main() {
  setStaticDir()

  http.HandleFunc("/", index)
  http.HandleFunc("/about", about)
  http.HandleFunc("/service", service)
  http.ListenAndServe(":8080", nil)
}
