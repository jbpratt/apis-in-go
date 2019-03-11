package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Value: "stranger",
			Usage: "your name here",
		},
		cli.IntFlag{
			Name:  "age",
			Value: 0,
			Usage: "your age here",
		},
	}

	// func parses and brings data in cli.Context struct
	app.Action = func(c *cli.Context) error {
		// c.String, c.Int looks for value of given flag
		log.Printf("Hello %s (%d), welcome", c.String("name"), c.Int("age"))
		return nil
	}
	// pass os.Args to cli app to parse
	app.Run(os.Args)
}
