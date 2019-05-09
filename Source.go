package main

import (
	"fmt"
	"github.com/urfave/cli"
)

func main() {
	fmt.Println("hello world")
	app := cli.NewApp()

	app.Name = "Vulnerous-GO, Website lookup in GOlang"

	app.Usage = "The app let's you\n 1) Query your IP\n2) CNAMES\n3) Mx Record query\n4) Query Name Servers\n"
}
