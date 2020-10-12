package cmd

import (
	"reflect"
	"strings"
)

// App for the main application
type App struct {
	args     []string
	msg      string
	exitCode int
}

var (
	// AvailableCmds for storing the available commands that can be used
	AvailableCmds = []string{"duplicate"}
)

// NewApp is for
func NewApp(args []string) *App {
	if len(args) > 0 {
		return &App{args: args[1:]}
	}

	return new(App)
}

// Log a string and set exit code as 0
func (a *App) Log(msg string) {
	a.msg = msg
	a.exitCode = 0
}

// LogError to report error
func (a *App) LogError(msg string, exitCode int) {
	a.msg = msg
	a.exitCode = exitCode
}

// Run the actual app
func (a *App) Run() (string, int) {
	if len(a.args) == 0 {
		a.Help()
	} else {
		cmdName := strings.ToLower(a.args[0])
		cmdMappings := map[string]func(){"duplicate": a.Duplicate}

		f := reflect.ValueOf(cmdMappings[cmdName])
		f.Call([]reflect.Value{})
	}

	return a.msg, a.exitCode
}
