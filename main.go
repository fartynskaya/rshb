package main

import (
  "fmt"
  "log"
  "net/http"
)

var responseText = "Hello, World"

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Add("Content-Type", "text/html; charset=utf-8")
  fmt.Fprint(w, responseText)
}
func main() {
  http.HandleFunc("/", HomeRouterHandler)
  log.Fatal(http.ListenAndServe(":8090", nil))
}