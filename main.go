package main

import (
  "fmt"
  //"log"
  "os"
  "github.com/codegangsta/cli"
  //"github.com/kureikain/gok"
)

func main() {
  fmt.Println("Welcome to Gok")

  app := cli.NewApp()
  app.Name = "Gok"
  app.Usage = "Hacker News bookmark with full text search in blevesearch"
  app.Action = func(c *cli.Context) {
    println("boom! I say!")
  }

  println(Version)

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
        error, _ := InitStorage("gok.db")
        if error == nil {
          println("Succesfully creae db")
        }
      },
    },
    {
      Name:      "add",
      ShortName: "a",
      Usage:     "add a link to the list",
      Action: func(c *cli.Context) {
        println("added :url ", c.Args().First())
        s := OpenStorage("gok.db")
        s.Add(c.Args().First())
      },
    },
    {
      Name:      "ls",
      ShortName: "l",
      Usage:     "list the link",
      Action: func(c *cli.Context) {
        s := OpenStorage("gok.db")
        s.List()
      },
    },
    {
      Name:      "date",
      ShortName: "d",
      Usage: "yesterday, today, last week, last month",
      Action: func(c *cli.Context) {
        println("added task: ", c.Args().First())
      },
    },
  }

  app.Run(os.Args)
}

