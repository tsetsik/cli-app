package cmd

import (
	"reflect"
	"strings"
)

// App for the main application
type App struct {
	args []string
}

var (
	// AvailableCmds for storing the available commands that can be used
	AvailableCmds = []string{"duplicate"}
)

// NewApp is for
func NewApp(args []string) *App {
	if len(args) > 0 {
		return &App{args[1:]}
	}

	return new(App)
}

// Run the actual app
func (a *App) Run() {
	if len(a.args) == 0 {
		a.Help()
	} else {
		cmdName := strings.ToLower(a.args[0])
		cmdMappings := map[string]func(){"duplicate": a.Duplicate}

		f := reflect.ValueOf(cmdMappings[cmdName])
		f.Call([]reflect.Value{})
	}
}
