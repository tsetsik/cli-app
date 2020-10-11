package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	app := &App{[]string{}}
	assert.Equal(t, NewApp([]string{}), app, "NewApp return something different than initiated App struct")
}

func TestRun(t *testing.T) {
	// When calling duplicate
	args := []string{"test", "duplicate"}
	app := NewApp(args)

	app.Run()
}
