package ggg

import "testing"

func TestAkk(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"silence is gold"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Akk()
		})
	}
}
