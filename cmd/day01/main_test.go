package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

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

func Test_readMeasurements(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "simple input",
			args: args{
				input: strings.NewReader("199\n200\n"),
			},
			want:    []int{199, 200},
			wantErr: false,
		},
		{
			name: "not a number",
			args: args{
				input: strings.NewReader("hello\nworld\n"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readMeasurements(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("readMeasurements() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readMeasurements() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateSlidingWindows(t *testing.T) {
	type args struct {
		measurements []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example from website",
			args: args{
				measurements: []int{
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
				},
			},
			want: []int{
				607,
				618,
				618,
				617,
				647,
				716,
				769,
				792,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateSlidingWindows(tt.args.measurements); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateSlidingWindows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_examplePart2(t *testing.T) {
	input := []int{
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
	}

	expectedResult := 5

	slidingWindows := generateSlidingWindows(input)
	result := countIncreases(slidingWindows)

	if result != expectedResult {
		t.Errorf("got = %v, want %v", result, expectedResult)
	}
}
