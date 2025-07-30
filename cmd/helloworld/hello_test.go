package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloTo(t *testing.T) {
	t.Run("say hello to the world", func(t *testing.T) {
		assert.Equal(t, "Hello world", Hello())
	})

	t.Run("say hello to the person", func(t *testing.T) {
		assert.Equal(t, "Hello, Gregory", HelloTo("Gregory"))
	})

	t.Run("say hello to the world when empty string is passed", func(t *testing.T) {
		actual := HelloTo("")
		assert.Equal(t, "Hello world", actual)
	})

	t.Run("say hallo in langugage passed as second argument", func(t *testing.T) {
		assert.Equal(t, "Witaj, Grzegorz", HelloTo("Grzegorz", "polish"))
		assert.Equal(t, "Hola, Gregory", HelloTo("Gregory", "spanish"))
		assert.Equal(t, "Bonjour, Gregory", HelloTo("Gregory", "french"))
	})
}
