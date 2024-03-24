package pkg

import "testing"

func Test_hello(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"void form pkg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Hellouuu()
		})
	}
}
