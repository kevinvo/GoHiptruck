package main

import (
  // "encoding/json"
  "fmt"
  // "io"
  // "io/ioutil"
  "net/http"
  // "github.com/gorilla/mux"
  "gohiptruck/models"
)

func Index(w http.ResponseWriter, r *http.Request) {  
  fmt.Fprintf(w, "Welcome!\n")
}

func Collection(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    http.Error(w, http.StatusText(405), 405)
    return
  }

  id := r.FormValue("id")
  if id == "" {
    http.Error(w, http.StatusText(400), 400)
    return
  }

  txn, err := models.GetTaxon(id)
  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }
  fmt.Fprintf(w, "%s %s", txn.Name, txn.Handle)
}
