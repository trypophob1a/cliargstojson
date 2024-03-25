package cliargstojson

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sum",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "sum",
			args: args{
				a: -1,
				b: 2,
			},
			want: 1,
		},
		{
			name: "sum",
			args: args{
				a: 3,
				b: -2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}