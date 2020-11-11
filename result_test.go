package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertSimpleBodyShouldReturnTrue(t *testing.T) {
	received := map[string]interface{}{
		"name": "eclésio melo",
		"age":  20,
	}

	expected := map[string]interface{}{
		"name": "eclésio melo",
		"age":  20,
	}

	assert.True(t, assertSimpleBody(received, expected))
}

func TestAssertArrayBodyShouldReturnTrue(t *testing.T) {
	received := []map[string]interface{}{
		{
			"name": "eclésio melo",
			"age":  20,
		},
	}

	expected := []map[string]interface{}{
		{
			"name": "eclésio melo",
			"age":  20,
		},
	}

	assert.True(t, assertArrayBody(received, expected))
}

func TestAssertArrayBodyShouldReturnFalse(t *testing.T) {
	received := []map[string]interface{}{
		{
			"name": "eclésio melo",
			"age":  20,
		},
	}

	expected := []map[string]interface{}{
		{
			"name": "eclésio melo",
			"age":  20,
		},
		{
			"name": "eclésio melo",
			"age":  20,
		},
	}

	assert.False(t, assertArrayBody(received, expected))
}
