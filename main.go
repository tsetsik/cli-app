package main

import (
	"cli-app/cmd"
	"os"
)

func main() {
	app := cmd.NewApp(os.Args)
	app.Run()
}
