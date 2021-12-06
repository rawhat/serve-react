package main

import (
  "fmt"
  "net/http"
)

func main() {
  mux := http.NewServeMux()
  mux.Handle("/", http.FileServer(http.Dir("./dist")))

  err := http.ListenAndServe(":8080", mux)
  if err != nil {
    fmt.Printf("Failed to listen and serve:  %v", err)
  }
}
