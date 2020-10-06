package main

import (
	"cli-app/cmd"
)

func main() {
	app := cmd.NewApp()
	app.Run()
}
