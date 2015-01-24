package main

import (
  "log"
  "github.com/boltdb/bolt"
  "fmt"
)

type Storage struct {
  Path string
}

var (
  db bolt.DB
)

func InitStorage(path string) (error, Storage) {
  s := Storage { Path: path, }

  // Open the my.db data file in your current directory.
  // It will be created if it doesn't exist.
  db, err := bolt.Open(s.Path, 0600, nil)
  if err != nil {
    return err, s
  }
  defer db.Close()

  db.Update(func(tx *bolt.Tx) error {
    _, err := tx.CreateBucket([]byte("MyList"))
    if err != nil {
      return fmt.Errorf("create bucket: %s", err)
    }
    return nil
  })

  return nil, s
}

func OpenStorage(path string) Storage {
  s := Storage { Path: path, }

  // Open the my.db data file in your current directory.
  // It will be created if it doesn't exist.
  db, err := bolt.Open(s.Path, 0600, nil)
  if err != nil {
      log.Fatal(err)
  }
  defer db.Close()

  return s
}

func (s Storage) Add(url string) bool {
  fmt.Println(url)
  OpenStorage(s.Path)

  db.Update(func(tx *bolt.Tx) error {

    b, err := tx.CreateBucketIfNotExists([]byte("MyList"))
    if err != nil {
        return err
    }
    result := b.Put([]byte(url), []byte("My New Year post"))

    return result

  })

  return true
}

func (s Storage) List() {
  OpenStorage(s.Path)

  err := db.View(func(tx *bolt.Tx) error {

    fmt.Println("view db")
    b := tx.Bucket([]byte("MyList"))
    c := b.Cursor()

    for k, v := c.First(); k != nil; k, v = c.Next() {
        fmt.Printf("key=%s, value=%s\n", k, v)
    }

    return nil
  })
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Lol")

}
