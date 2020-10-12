package cmd

// Duplicate func for finding a duplicate in a given slice
func (a *App) Duplicate() {
	if len(a.args) <= 1 {
		a.LogError("You need to provide argument for search eg. `go run main.go duplicate [1, 2, 3]`", 1)
	} else {
		a.Log("Here are the duplicates")
	}
}
