package main

import (
	"os"
	"testing"
  "fmt"
	"github.com/boltdb/bolt"
)

var (
	s Storage
)

func init() {
	s, _ = NewStorage("/tmp/gok")
  cleanup()
}

func cleanup() {
  s.DB.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte("MyList"))
		if err != nil {
			return nil
		}
    return nil
	})
}

func TestCreateNewStorage(t *testing.T) {

	if s.Path != "/tmp/gok" {
		t.Error("Wrong storage path")
	}

	if _, err := os.Stat("/tmp/gok.bleve"); os.IsNotExist(err) {
		t.Error("Bleve index is not created")
	}

	if _, err := os.Stat("/tmp/gok.bolt"); os.IsNotExist(err) {
		t.Error("Bolt db is not created")
	}
}

func TestAddNewItem(t *testing.T) {
	item, err := NewItem("http://go-wise.blogspot.com/2011/10/running-tests-in-parallel.html")
	if err != nil {
		t.Error("Cannot add item")
	}
	s.Add(item)
}

func TestListItem(t *testing.T) {
	result, err := s.List()
  if err != nil {
    t.Error("Cannot fetch the list")
  }

  if len(result) <1 {
    t.Error("Empty result")
  }

  if result[0].Url != "http://go-wise.blogspot.com/2011/10/running-tests-in-parallel.html" {
    t.Error("Incorrect item order")
  }
}

func TestSearchItem(t *testing.T) {
	item, _ := NewItem("http://blog.steventroughtonsmith.com/post/109040361205/mpw-carbon-and-building-classic-mac-os-apps-in-os")
  s.Add(item)

  result, error := s.Search("Test")

  if error != nil {
    t.Error(fmt.Println(error))
  }

  if len(result) < 2 {
    t.Error("Not found all")
  }

  if result[0].Url != "http://go-wise.blogspot.com/2011/10/running-tests-in-parallel.html" {
    t.Error("Incorrect item order " , result[0].Url)
  }
  if result[1].Url != "http://blog.steventroughtonsmith.com/post/109040361205/mpw-carbon-and-building-classic-mac-os-apps-in-os" {
    t.Error("Incorrect item order ", result[1].Url)
  }
}
