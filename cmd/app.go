package cmd

import (
	"fmt"
)

// App for the main application
type App struct {
}

// NewApp is for
func NewApp() *App {
	return new(App)
}

// Run the actual app
func (a *App) Run() {
	fmt.Println("\n\n What we're going to see here")
}
