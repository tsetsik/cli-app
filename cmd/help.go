package cmd

import (
	"strings"
)

// Help to print whenever a command is not supplied
func (a *App) Help() {
	a.Log("The available commands are: " + strings.Join(AvailableCmds, ","))
}
