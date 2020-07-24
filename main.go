package main

import (
	"fmt"
	"log"
	"os"

	"github.com/KeThichDua/ex5go/reqres5"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

// Commands cli
func Commands() {
	app.Commands = []cli.Command{
		{
			Name:  "Run2",
			Usage: "Run request 2 of exercise 5",
			Action: func(c *cli.Context) {
				Run2()
			},
		},
		{
			Name:  "Run3",
			Usage: "Run request 3 of exercise 5",
			Action: func(c *cli.Context) {
				Run3()
			},
		},
		{
			Name:  "Run4",
			Usage: "Run request 4 of exercise 5",
			Action: func(c *cli.Context) {
				Run4()
			},
		},
		{
			Name:  "Run5",
			Usage: "Run request 5 of exercise 5",
			Action: func(c *cli.Context) {
				reqres5.Run5()
			},
		},
	}
}

func main() {
	// Run2()
	// Run3()
	// Run4()
	// reqres5.Run5()
	Commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("done")
}
