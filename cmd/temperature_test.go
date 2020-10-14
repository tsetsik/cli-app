package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemperature(t *testing.T) {
	// With -15 should return -20
	args := []string{"test", "temperature", "-15", "[3, -20, 1, 8, 7, 4, -3]"}
	app := NewApp(args)

	exitCode := app.Temperature()

	assert.Equal(t, exitCode, 0)
	assert.Equal(t, app.msg, "The answer is: -20")

	// With 6 should return 7
	args = []string{"test", "temperature", "6", "[3, -20, 1, 8, 7, 4, -3]"}
	app = NewApp(args)

	exitCode = app.Temperature()

	assert.Equal(t, exitCode, 0)
	assert.Equal(t, app.msg, "The answer is: 7")

	// With 1 should return 0
	args = []string{"test", "temperature", "1", "[2, 0, 5]"}
	app = NewApp(args)

	exitCode = app.Temperature()

	assert.Equal(t, exitCode, 0)
	assert.Equal(t, app.msg, "The answer is: 0")
}
