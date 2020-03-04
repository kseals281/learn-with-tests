package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		language string
		want     string
	}{
		{
			"to a person",
			"Chris",
			"",
			"Hello, Chris",
		}, {
			"empty string",
			"",
			"",
			"Hello, World",
		}, {
			"in Spanish",
			"Elodie",
			"Spanish",
			"Hola, Elodie",
		}, {
			"in French",
			"Lauren",
			"French",
			"Bonjour, Lauren",
		},
	}

	for _, tt := range tests {
		got := Hello(tt.name, tt.language)
		if got != tt.want {
			t.Errorf("got %q want %q", got, tt.want)
		}
	}

}
