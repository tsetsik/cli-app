package cmd

import (
	"fmt"
	"testing"
)

func TestDuplicate(t *testing.T) {
	// With duplicates
	args := []string{"test", "duplicate", "[1, 2, 1, 3, 4, 5, 3]"}
	app := NewApp(args)

	app.Duplicate()

	fmt.Println("\n\n Test da go evaaaaaaa: ", app.msg)
}
