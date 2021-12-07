package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_readCrabPositions(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "simple test",
			args: args{
				reader: strings.NewReader("1,2"),
			},
			want:    []int{1, 2},
			wantErr: false,
		},
		{
			name: "website example",
			args: args{
				reader: strings.NewReader("16,1,2,0,4,2,7,1,2,14"),
			},
			want:    []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readCrabPositions(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("readCrabPositions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readCrabPositions() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMinimum(t *testing.T) {
	type args struct {
		positions []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ordered",
			args: args{
				positions: []int{1, 2},
			},
			want: 1,
		}, {
			name: "unordered",
			args: args{
				positions: []int{5, 7, 3, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinimum(tt.args.positions); got != tt.want {
				t.Errorf("getMinimum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMaximum(t *testing.T) {
	type args struct {
		positions []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ordered",
			args: args{
				positions: []int{1, 2},
			},
			want: 2,
		},
		{
			name: "unordered",
			args: args{
				positions: []int{5, 7, 3, 2},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximum(tt.args.positions); got != tt.want {
				t.Errorf("getMaximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateCost(t *testing.T) {
	type args struct {
		positions      []int
		targetPosition int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple test",
			args: args{
				positions:      []int{1, 2},
				targetPosition: 1,
			},
			want: 1,
		},
		{
			name: "website example",
			args: args{
				positions:      []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
				targetPosition: 2,
			},
			want: 37,
		},
		{
			name: "website example 2",
			args: args{
				positions:      []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
				targetPosition: 3,
			},
			want: 39,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateCost(tt.args.positions, tt.args.targetPosition); got != tt.want {
				t.Errorf("calculateCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMinimumCost_part1(t *testing.T) {
	type args struct {
		positions       []int
		minimumPosition int
		maximumPosition int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple test",
			args: args{
				positions: []int{1, 2},
			},
			want: 1,
		},
		{
			name: "website example",
			args: args{
				positions: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			},
			want: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinimumCost(tt.args.positions, calculateCost); got != tt.want {
				t.Errorf("findMinimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateCostsCorrectly(t *testing.T) {
	type args struct {
		positions      []int
		targetPosition int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple test",
			args: args{
				positions:      []int{1, 2},
				targetPosition: 1,
			},
			want: 1,
		},
		{
			name: "website example",
			args: args{
				positions:      []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
				targetPosition: 2,
			},
			want: 206,
		},
		{
			name: "website example 2",
			args: args{
				positions:      []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
				targetPosition: 5,
			},
			want: 168,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateCostsCorrectly(tt.args.positions, tt.args.targetPosition); got != tt.want {
				t.Errorf("calculateCostsCorrectly() = %v, want %v", got, tt.want)
			}
		})
	}
}
