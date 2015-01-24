package main

import (
  "testing"
  "os"
)

func TestCreateNewStorage(t *testing.T) {
  s, err := NewStorage("/tmp/gok")
  
  if s.Path != "/tmp/gok" {
    t.Error("Wrong storage path")
  }

  if err != nil {
    t.Error("Storage fail to create")
  }

  if _, err := os.Stat("/tmp/gok.bleve"); os.IsNotExist(err) {
      t.Error( "Bleve index is not created")
  }

  if _, err := os.Stat("/tmp/gok.bolt"); os.IsNotExist(err) {
      t.Error( "Bolt db is not created")
  }
}

func TestAddNewItem(t *testing.T) {

}

func TestListItem(t *testing.T) {

}

func TestSearchItem(t *testing.T) {
}

