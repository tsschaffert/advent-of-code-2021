package main

import "testing"

func Test_countIncreases(t *testing.T) {
	type args struct {
		measurements []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "example from website",
			args: struct{ measurements []int }{measurements: []int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			}},
			want: 7,
		},
		{
			name: "simple example",
			args: struct{ measurements []int }{measurements: []int{
				199,
				200,
			}},
			want: 1,
		},
		{
			name: "only one measurement",
			args: struct{ measurements []int }{measurements: []int{
				199,
			}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countIncreases(tt.args.measurements); got != tt.want {
				t.Errorf("countIncreases() = %v, want %v", got, tt.want)
			}
		})
	}
}
