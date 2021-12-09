package main

import (
	"io"
	"math"
	"reflect"
	"strings"
	"testing"
)

func Test_readHeightmap(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    Heightmap
		wantErr bool
	}{
		{
			name: "simple example",
			args: args{
				input: strings.NewReader("12\n34\n"),
			},
			want: Heightmap{
				[]int{1, 2},
				[]int{3, 4},
			},
			wantErr: false,
		},
		{
			name: "website example",
			args: args{
				input: strings.NewReader("2199943210\n3987894921\n9856789892\n8767896789\n9899965678\n"),
			},
			want: Heightmap{
				[]int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
				[]int{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
				[]int{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
				[]int{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
				[]int{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readHeightmap(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("readHeightmap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readHeightmap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLowpoints(t *testing.T) {
	type args struct {
		heightmap Heightmap
	}
	tests := []struct {
		name string
		args args
		want []Point
	}{
		{
			name: "simple example",
			args: args{
				heightmap: Heightmap{
					[]int{1, 2},
					[]int{3, 4},
				},
			},
			want: []Point{{x: 0, y: 0}},
		},
		{
			name: "website example",
			args: args{
				heightmap: Heightmap{
					[]int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
					[]int{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
					[]int{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
					[]int{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
					[]int{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
				},
			},
			want: []Point{
				{x: 0, y: 1},
				{x: 0, y: 9},
				{x: 2, y: 2},
				{x: 4, y: 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLowpoints(tt.args.heightmap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLowpoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeightmap_getHeightForLowpoints(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		hm   Heightmap
		args args
		want int
	}{
		{
			name: "simple example",
			hm: Heightmap{
				[]int{1, 2},
				[]int{3, 4},
			},
			args: args{x: 0, y: 0},
			want: 1,
		},
		{
			name: "x not existing example",
			hm: Heightmap{
				[]int{1, 2},
				[]int{3, 4},
			},
			args: args{x: 2, y: 0},
			want: math.MaxInt,
		},
		{
			name: "y not existing example",
			hm: Heightmap{
				[]int{1, 2},
				[]int{3, 4},
			},
			args: args{x: 0, y: -1},
			want: math.MaxInt,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hm.getHeightForLowpoints(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("getHeightForLowpoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateRiskLevel(t *testing.T) {
	type args struct {
		height int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple test",
			args: args{height: 0},
			want: 1,
		},
		{
			name: "very complex test",
			args: args{height: 8},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateRiskLevel(tt.args.height); got != tt.want {
				t.Errorf("calculateRiskLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateSumOfRiskLevels(t *testing.T) {
	type args struct {
		heightmap Heightmap
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple example",
			args: args{
				heightmap: Heightmap{
					[]int{1, 2},
					[]int{3, 4},
				},
			},
			want: 2,
		},
		{
			name: "website example",
			args: args{
				heightmap: Heightmap{
					[]int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
					[]int{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
					[]int{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
					[]int{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
					[]int{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
				},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSumOfRiskLevels(tt.args.heightmap); got != tt.want {
				t.Errorf("calculateSumOfRiskLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_detectBasin(t *testing.T) {
	type args struct {
		heightmap Heightmap
		lowPoint  Point
	}
	tests := []struct {
		name string
		args args
		want []Point
	}{
		{
			name: "single point test",
			args: args{
				heightmap: Heightmap{
					[]int{1, 9},
					[]int{9, 9},
				},
				lowPoint: Point{x: 0, y: 0},
			},
			want: []Point{{x: 0, y: 0}},
		},
		{
			name: "basic test",
			args: args{
				heightmap: Heightmap{
					[]int{1, 3},
					[]int{4, 9},
				},
				lowPoint: Point{x: 0, y: 0},
			},
			// TODO order should not matter
			want: []Point{
				{x: 0, y: 0},
				{x: 1, y: 0},
				{x: 0, y: 1},
			},
		},
		{
			name: "website example",
			args: args{
				heightmap: Heightmap{
					[]int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
					[]int{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
					[]int{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
					[]int{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
					[]int{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
				},
				lowPoint: Point{x: 0, y: 1},
			},
			// TODO order should not matter
			want: []Point{
				{x: 0, y: 1},
				{x: 0, y: 0},
				{x: 1, y: 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectBasin(tt.args.heightmap, tt.args.lowPoint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectBasin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getProductOfTop3BasinSizes(t *testing.T) {
	type args struct {
		heightmap Heightmap
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "website example",
			args: args{
				heightmap: Heightmap{
					[]int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
					[]int{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
					[]int{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
					[]int{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
					[]int{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
				},
			},
			want: 1134,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getProductOfTop3BasinSizes(tt.args.heightmap); got != tt.want {
				t.Errorf("getProductOfTop3BasinSizes() = %v, want %v", got, tt.want)
			}
		})
	}
}
