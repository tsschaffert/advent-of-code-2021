package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_readPopulation(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []lanternfish
		wantErr bool
	}{
		{
			name: "simple test",
			args: args{
				input: strings.NewReader("1,2"),
			},
			want:    []lanternfish{{1}, {2}},
			wantErr: false,
		},
		{
			name: "website example",
			args: args{
				input: strings.NewReader("3,4,3,1,2"),
			},
			want:    []lanternfish{{3}, {4}, {3}, {1}, {2}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readPopulation(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("readPopulation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readPopulation() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lanternfish_simulate(t *testing.T) {
	fish := lanternfish{1}
	expectedFish := lanternfish{0}

	fish.simulate()

	if !reflect.DeepEqual(fish, expectedFish) {
		t.Errorf("lanternfish timer got = %v, want %v", fish, expectedFish)
	}
}

func Test_lanternfish_simulate_overflow(t *testing.T) {
	fish := lanternfish{1}
	expectedFish := lanternfish{6}

	fish.simulate()
	fish.simulate()

	if !reflect.DeepEqual(fish, expectedFish) {
		t.Errorf("lanternfish timer got = %v, want %v", fish, expectedFish)
	}
}

func Test_lanternfish_simulate_return_value(t *testing.T) {
	fish := lanternfish{1}
	expectedSpawnNewFish := false

	spawnNewFish := fish.simulate()

	if spawnNewFish != expectedSpawnNewFish {
		t.Errorf("lanternfish.simulate() got = %v, want %v", spawnNewFish, expectedSpawnNewFish)
	}
}

func Test_lanternfish_simulate_overflow_return_value(t *testing.T) {
	fish := lanternfish{1}
	expectedSpawnNewFish := true

	fish.simulate()
	spawnNewFish := fish.simulate()

	if spawnNewFish != expectedSpawnNewFish {
		t.Errorf("lanternfish.simulate() got = %v, want %v", spawnNewFish, expectedSpawnNewFish)
	}
}

func Test_simulatePopulationStep(t *testing.T) {
	type args struct {
		population []lanternfish
	}
	tests := []struct {
		name string
		args args
		want []lanternfish
	}{
		{
			name: "website example day 1",
			args: args{
				population: []lanternfish{{3}, {4}, {3}, {1}, {2}},
			},
			want: []lanternfish{{2}, {3}, {2}, {0}, {1}},
		},
		{
			name: "website example day 2",
			args: args{
				population: []lanternfish{{2}, {3}, {2}, {0}, {1}},
			},
			want: []lanternfish{{1}, {2}, {1}, {6}, {0}, {8}},
		},
		{
			name: "website example with newly spawned fish",
			args: args{
				population: []lanternfish{{1}, {2}, {1}, {6}, {0}, {8}},
			},
			want: []lanternfish{{0}, {1}, {0}, {5}, {6}, {7}, {8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simulatePopulationStep(tt.args.population); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simulatePopulationStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_simulatePopulation(t *testing.T) {
	type args struct {
		population []lanternfish
		steps      int
	}
	tests := []struct {
		name string
		args args
		want []lanternfish
	}{
		{
			name: "website example",
			args: args{
				population: []lanternfish{{3}, {4}, {3}, {1}, {2}},
				steps:      2,
			},
			want: []lanternfish{{1}, {2}, {1}, {6}, {0}, {8}},
		},
		{
			name: "website example 3 days",
			args: args{
				population: []lanternfish{{3}, {4}, {3}, {1}, {2}},
				steps:      3,
			},
			want: []lanternfish{{0}, {1}, {0}, {5}, {6}, {7}, {8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simulatePopulation(tt.args.population, tt.args.steps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simulatePopulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
