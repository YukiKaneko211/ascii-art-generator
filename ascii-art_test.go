package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name string
		arg1 string
	}{
		{"", ""},
		{"Hello\n", "Hello\n"},
		{"hello", "hello"},
		{"HeLlO", "HeLlO"},
		{"Hello There", "Hello There"},
		{"1Hello 2There", "1Hello 2There"},
		{"{Hello There}", "{Hello There}"},
		{"Hello\nThere", "Hello\nThere"},
		{"Hello\n\nThere", "Hello\n\nThere"},
		{"\n", "\n"},
		{"HELLO", "HELLO"},
		{"HeLlo HuMaN", "HeLlo HuMaN"},
		{"{Hello & There #}", "{Hello & There #}"},
		{"MaD3IrA&LiSboN", "MaD3IrA&LiSboN"},
		{"1a\"#FdwHywR&/()=", "1a\"#FdwHywR&/()="},
		{"{|}~", "{|}~"},
		{"RGB", "RGB"},
		{":;<=>?@", ":;<=>?@"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "ABCDEFGHIJKLMNOPQRSTUVWXYZ"},
		{"abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save the original command-line arguments
			originalArgs := os.Args
			defer func() {
				// Restore the original command-line arguments when the test is done
				os.Args = originalArgs
			}()

			// Set up new command-line arguments
			os.Args = []string{"cmd", tt.arg1}

			// Call the main function
			main()
		})
	}
}
