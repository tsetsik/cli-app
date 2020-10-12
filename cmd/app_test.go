package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	app := &App{}
	assert.Equal(t, NewApp([]string{}), app, "NewApp return something different than initiated App struct")
}

func TestRun(t *testing.T) {
	// When calling a command that returns code 0
	args := []string{"test", "duplicate", "[1]"}
	app := NewApp(args)

	msg, exitCode := app.Run()

	assert.NotEmpty(t, msg)
	assert.Equal(t, exitCode, 0)

	// When calling a command that returns code 1
	args = []string{"test", "duplicate"}
	app = NewApp(args)

	msg, exitCode = app.Run()

	assert.NotEmpty(t, msg)
	assert.Equal(t, exitCode, 1)
}
