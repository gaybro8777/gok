package main

import (
	"strings"
	"testing"
)

func TestNewItem(t *testing.T) {
	url := "http://go-wise.blogspot.com/2011/10/running-tests-in-parallel.html"
	item, err := NewItem(url)

	if err != nil {
		t.Error("Fail to create URL ITEM")
	}

	if item.Url != url {
		t.Error("Creating wrong")
	}
	if item.Title != "Adentures in Go: Running Tests In Parallel" {
		t.Error("Fail to parse Title")
	}

	if strings.Index(item.Body, "Running Tests In Parallel") == -1 {
		t.Error("Fail to parse Body")
	}

}
