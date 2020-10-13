package cmd

import (
	"encoding/json"
)

// Duplicate func for finding a duplicate in a given slice
func (a *App) Duplicate() int {
	if len(a.args) <= 1 {
		a.Log("You need to provide argument for search eg. `go run main.go duplicate [1, 2, 3]`")
		return 1
	}

	var foo []int
	f := []byte(a.args[1])
	if err := json.Unmarshal(f, &foo); err != nil {
		a.Log("You provided not supported argument")
		return 1
	}

	a.Log("Here are the duplicates")
	return 0
}
