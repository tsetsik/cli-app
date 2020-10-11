package cmd

import (
	"fmt"
)

// App for the main application
type App struct {
	args []string
}

// NewApp is for
func NewApp(args []string) *App {
	return &App{args[1:]}
}

// Run the actual app
func (a *App) Run() {
	fmt.Println("\n\n What we're going to see here: ", a.args)
}
