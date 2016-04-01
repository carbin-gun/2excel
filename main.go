package main

import (
	"os"

	"fmt"

	"regexp"

	"github.com/codegangsta/cli"
)

const Version = "1.0.0"

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Author = "carbin-gun"
	app.Email = "cilendeng@gmail.com"
	app.Name = "2excel"
	app.Usage = "convert sql export or csv file to excel with one command!"
	app.Action = func(c *cli.Context) {
		command := parse(c)
		DoConvert(command)
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "f",
			Usage: "the file which the program will convert from",
		},
		cli.StringFlag{
			Name:  "d",
			Usage: "the delimeter of different fields of the same row.the tab or comma is support by default,you can point it out or not",
		},
		cli.StringFlag{
			Name:  "t",
			Usage: "the target directory the excel will be generated at,the default is current dir.",
		},
	}

	app.Run(os.Args)
}

func parse(c *cli.Context) Command {
	source := c.String("f")
	targetDir := c.String("t")
	delimiter := c.String("d")
	if delimiter != "" {
		delimiter = regexp.QuoteMeta(delimiter)
	}
	fmt.Println("args:", c.Args())
	if source == "" {
		if len(c.Args()) != 1 {
			panic("please provide the file you need to process. ")
		}
		source = c.Args()[0]
	}

	command := Command{
		SourceFile: source,
		Delimiter:  delimiter,
		Dest:       targetDir,
	}
	return command

}
