package main

import (
	"cli-app/cmd"
	"os"
)

func main() {
	app := cmd.NewApp(os.Args)
	msg, exitCode := app.Run()
	os.Stdout.WriteString(msg)
	os.Exit(exitCode)
}
