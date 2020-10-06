package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	assert.Equal(t, NewApp(), new(App), "NewApp return something different than initiated App struct")
}
