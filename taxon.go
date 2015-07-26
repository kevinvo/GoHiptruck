package main
// package db

import (
  "fmt"
  "database/sql"
  "log"
)

type Taxon struct {
  Name string
  Handle string
}

func GetTaxon(id string) (Taxon, error){
  txn := new(Taxon)
  query := "SELECT name, handle FROM catalog_taxons WHERE id = $1"
  err := db.QueryRow(query, id).
            Scan(&txn.Name, &txn.Handle)
  switch {
  case err == sql.ErrNoRows:
          log.Printf("No user with that ID.")
  case err != nil:
          log.Fatal(err)
  default:
          fmt.Printf("txn is %s\n", txn)
  }
  return *txn, nil
}