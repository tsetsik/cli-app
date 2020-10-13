package cmd

import (
	"reflect"
	"strings"
)

// App for the main application
type App struct {
	args []string
	msg  string
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
}

// Run the actual app
func (a *App) Run() (string, int) {
	exitCode := 0
	if len(a.args) == 0 {
		a.Help()
	} else {
		cmdName := strings.ToLower(a.args[0])
		cmdMappings := map[string]func() int{"duplicate": a.Duplicate}

		f := reflect.ValueOf(cmdMappings[cmdName])
		exitCode = f.Call([]reflect.Value{})[0].Interface().(int)
	}

	return a.msg, exitCode
}
