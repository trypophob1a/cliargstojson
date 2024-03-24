package pkg

import "testing"

func TestFFF(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"void"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FFF()
		})
	}
}
