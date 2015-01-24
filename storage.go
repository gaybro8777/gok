package main

import (
  "log"
  "github.com/boltdb/bolt"
  "fmt"
)

type Storage struct {
  Path string
  DB   *bolt.DB
}

var (
  db *bolt.DB
)

func NewStorage(path string) (Storage, error) {
  s := Storage { Path: path, }
  db, err := bolt.Open(s.Path, 0600, nil)
  if err != nil {
    return s, err
  }
  log.Println(db)
  s.DB = db
  //defer db.Close()

  return s, nil
}

func (s Storage) Add(item *Item) bool {
  s.DB.Update(func(tx *bolt.Tx) error {

    b, err := tx.CreateBucketIfNotExists([]byte("MyList"))
    if err != nil {
        return err
    }
    result := b.Put([]byte(item.Url), []byte(item.Title))
    return result
  })

  return true
}

func (s Storage) List() {
  err := s.DB.View(func(tx *bolt.Tx) error {

    b := tx.Bucket([]byte("MyList"))
    c := b.Cursor()

    fmt.Println("-----------------------------------------")
    fmt.Println("|Key----------------|Value------------------")
    for k, v := c.First(); k != nil; k, v = c.Next() {
        fmt.Printf("|%s | %s\n|", k, v)
    }
    fmt.Println("----------------------------------------")

    return nil
  })
  if err != nil {
    log.Fatal(err)
  }
}
