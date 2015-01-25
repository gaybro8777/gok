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
	item, err := NewItem("http://google.com")
	if err != nil {
		t.Error("Cannot add item")
	}
	s.Add(item)
}

func TestListItem(t *testing.T) {
	s.List()
}

func TestSearchItem(t *testing.T) {
	s.Search("Google")
}
