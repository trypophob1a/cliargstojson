package ggg

import "testing"

func TestGgg(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"void"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Wed()
		})
	}
}