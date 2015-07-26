package main

import (
  "log"
  "net/http"
)

func main() {
  InitDB("user=kevinvo1 dbname=hipvan_staging sslmode=disable")
  router := NewRouter()
  log.Fatal(http.ListenAndServe(":8080", router))
}
