package main

import (
  "testing"
)

func TestNewItem(t *testing.T) {
  item := &Item{ "http://github.com/", "", ""}

  if item.Url != "http://github.com/" {
    t.Error("Creating wrong")
  }
  
  if item.Title != "GitHub" {
    t.Error("Fail to fetch URL")
  }

}

