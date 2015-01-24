package main

import (
  "log"
  "github.com/boltdb/bolt"
  "fmt"

  "github.com/blevesearch/bleve"
)

type Storage struct {
  Path string
  DB   *bolt.DB
  Index bleve.Index
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

  index, err := bleve.Open("gok.bleve")
  if err == bleve.ErrorIndexPathDoesNotExist {
    log.Printf("Creating new index...")
    // create a mapping
    indexMapping, err := buildIndexMapping()
    if err != nil {
      log.Fatal(err)
    }
    index, err = bleve.New("gok.bleve", indexMapping)
    if err != nil {
      log.Fatal(err)
    }
  }

  s.Index = index

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

  s.Index.Index(item.Url, item)

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

func (s Storage) Search(url string) {
  // search for some text
    query := bleve.NewMatchQuery(url)
    search := bleve.NewSearchRequest(query)
    searchResults, err := s.Index.Search(search)
    if err != nil {
        fmt.Println(err)
        return
    }
  fmt.Println(searchResults)
}

func buildIndexMapping() (*bleve.IndexMapping, error) {
  englishTextFieldMapping := bleve.NewTextFieldMapping()
  englishTextFieldMapping.Analyzer = "en"

  // a generic reusable mapping for keyword text
  keywordFieldMapping := bleve.NewTextFieldMapping()
  keywordFieldMapping.Analyzer = "keyword"

  urlMapping := bleve.NewDocumentMapping()

  urlMapping.AddFieldMappingsAt("url", keywordFieldMapping)

  urlMapping.AddFieldMappingsAt("title", englishTextFieldMapping)
  urlMapping.AddFieldMappingsAt("body", keywordFieldMapping)

  indexMapping := bleve.NewIndexMapping()
  indexMapping.AddDocumentMapping("url", urlMapping)
  indexMapping.TypeField = "type"
  indexMapping.DefaultAnalyzer = "en"
  return indexMapping, nil
}
