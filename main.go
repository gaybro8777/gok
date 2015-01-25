package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	//"github.com/kureikain/gok"
)

func main() {

	app := cli.NewApp()
	app.Name = "Gok"
	app.Usage = "A bookmark with full text search in blevesearch and bolt db"
	app.Action = func(c *cli.Context) {
		fmt.Println("Welcome to " + name)
	}

	app.Commands = []cli.Command{
		{
			Name:      "version",
			ShortName: "v",
			Usage:     "Version",
			Action: func(c *cli.Context) {
				println("Gok v ", Version)
			},
		},
		{
			Name:      "init",
			ShortName: "i",
			Usage:     "add a link to the list",
			Action: func(c *cli.Context) {
				println("added task: ", c.Args().First())
				_, err := NewStorage("gok")
				if err != nil {
					println("Succesfully creae db")
				}
			},
		},
		{
			Name:      "add",
			ShortName: "a",
			Usage:     "add a link to the list",
			Action: func(c *cli.Context) {
				s, _ := NewStorage("gok")
				item, err := NewItem(c.Args().First())

				if err != nil {
					log.Fatal(err)
				}
				if item == nil {
					log.Fatal("Invalid URL")
				}

				s.Add(item)
			},
		},
		{
			Name:      "ls",
			ShortName: "l",
			Usage:     "list the link",
			Action: func(c *cli.Context) {
				s, _ := NewStorage("gok")
				result, _ := s.List()
        for key, value := range result {
          fmt.Printf("%d. %s = %s\n", key, value.Url, value.Title)
        }
			},
		},
		{
			Name:      "date",
			ShortName: "d",
			Usage:     "yesterday, today, last week, last month",
			Action: func(c *cli.Context) {
				println("added task: ", c.Args().First())
			},
		},
		{
			Name:      "search",
			ShortName: "s",
			Usage:     "Search link by keyword",
			Action: func(c *cli.Context) {
				s, _ := NewStorage("gok")
				s.Search(c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}
