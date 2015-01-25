package main

import (
	"os"
	"testing"
)

var (
	s Storage
)

func init() {
	s, _ = NewStorage("/tmp/gok")
	/*if err != nil {*/
	/*t.Error("Storage fail to create")*/
	/*}*/
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

/*func TestSearchItem(t *testing.T) {*/
	/*s.Search("Google")*/
/*}*/
