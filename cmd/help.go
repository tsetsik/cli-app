package cmd

import (
	"os"
	"strings"
)

// Help to print whenever a command is not supplied
func (a *App) Help() {
	os.Stdout.WriteString("The available commands are: " + strings.Join(AvailableCmds, ","))
}
