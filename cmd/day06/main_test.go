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

func Test_population_simulate(t *testing.T) {
	p := densePopulation{buckets: [9]int64{0, 1, 1, 2, 1, 0, 0, 0, 0}}
	expectedPopulation := densePopulation{buckets: [9]int64{1, 1, 2, 1, 0, 0, 0, 0, 0}}

	p.simulate()

	if p != expectedPopulation {
		t.Errorf("lanternfish.simulate() got = %v, want %v", p, expectedPopulation)
	}
}

func Test_population_simulate_overflow(t *testing.T) {
	p := densePopulation{buckets: [9]int64{1, 1, 2, 1, 0, 0, 0, 0, 0}}
	expectedPopulation := densePopulation{buckets: [9]int64{1, 2, 1, 0, 0, 0, 1, 0, 1}}

	p.simulate()

	if p != expectedPopulation {
		t.Errorf("lanternfish.simulate() got = %v, want %v", p, expectedPopulation)
	}
}

func Test_population_simulate_overflow_with_existing_fish_at_6(t *testing.T) {
	p := densePopulation{buckets: [9]int64{1, 1, 2, 1, 0, 0, 0, 1, 0}}
	expectedPopulation := densePopulation{buckets: [9]int64{1, 2, 1, 0, 0, 0, 2, 0, 1}}

	p.simulate()

	if p != expectedPopulation {
		t.Errorf("lanternfish.simulate() got = %v, want %v", p, expectedPopulation)
	}
}

func Test_convertToDensePopulation(t *testing.T) {
	type args struct {
		population []lanternfish
	}
	tests := []struct {
		name string
		args args
		want densePopulation
	}{
		{
			name: "simple example",
			args: args{
				population: []lanternfish{{1}, {2}},
			},
			want: densePopulation{buckets: [9]int64{0, 1, 1, 0, 0, 0, 0, 0, 0}},
		},
		{
			name: "simple example",
			args: args{
				population: []lanternfish{{3}, {4}, {3}, {1}, {2}},
			},
			want: densePopulation{buckets: [9]int64{0, 1, 1, 2, 1, 0, 0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToDensePopulation(tt.args.population); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToDensePopulation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_simulateDensePopulation(t *testing.T) {
	type args struct {
		population densePopulation
		steps      int
	}
	tests := []struct {
		name string
		args args
		want densePopulation
	}{
		{
			name: "simple test",
			args: args{
				population: densePopulation{buckets: [9]int64{0, 1, 1, 0, 0, 0, 0, 0, 0}},
				steps:      1,
			},
			want: densePopulation{buckets: [9]int64{1, 1, 0, 0, 0, 0, 0, 0, 0}},
		},
		{
			name: "website example 2 days",
			args: args{
				population: densePopulation{buckets: [9]int64{0, 1, 1, 2, 1, 0, 0, 0, 0}},
				steps:      2,
			},
			want: densePopulation{buckets: [9]int64{1, 2, 1, 0, 0, 0, 1, 0, 1}},
		},
		{
			name: "website example 18 days",
			args: args{
				population: densePopulation{buckets: [9]int64{0, 1, 1, 2, 1, 0, 0, 0, 0}},
				steps:      18,
			},
			want: densePopulation{buckets: [9]int64{3, 5, 3, 2, 2, 1, 5, 1, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simulateDensePopulation(tt.args.population, tt.args.steps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simulateDensePopulation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_densePopulation_size(t *testing.T) {
	type fields struct {
		buckets [9]int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "simple test",
			fields: fields{
				buckets: [9]int64{0, 1, 1, 0, 0, 0, 0, 0, 0},
			},
			want: 2,
		},
		{
			name: "website example",
			fields: fields{
				buckets: [9]int64{0, 1, 1, 2, 1, 0, 0, 0, 0},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := densePopulation{
				buckets: tt.fields.buckets,
			}
			if got := p.size(); got != tt.want {
				t.Errorf("size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_websiteExample_256(t *testing.T) {
	p := densePopulation{buckets: [9]int64{0, 1, 1, 2, 1, 0, 0, 0, 0}}
	var expectedPopulationCount int64 = 26984457539

	p = simulateDensePopulation(p, 256)
	populationCount := p.size()

	if populationCount != expectedPopulationCount {
		t.Errorf("website example population count after 256 iterations got = %v, want %v", populationCount, expectedPopulationCount)
	}
}

func Test_websiteExample_80(t *testing.T) {
	p := densePopulation{buckets: [9]int64{0, 1, 1, 2, 1, 0, 0, 0, 0}}
	var expectedPopulationCount int64 = 5934

	p = simulateDensePopulation(p, 80)
	populationCount := p.size()

	if populationCount != expectedPopulationCount {
		t.Errorf("website example population count after 80 iterations got = %v, want %v", populationCount, expectedPopulationCount)
	}
}

func Test_websiteExample_18(t *testing.T) {
	p := densePopulation{buckets: [9]int64{0, 1, 1, 2, 1, 0, 0, 0, 0}}
	var expectedPopulationCount int64 = 26

	p = simulateDensePopulation(p, 18)
	populationCount := p.size()

	if populationCount != expectedPopulationCount {
		t.Errorf("website example population count after 18 iterations got = %v, want %v", populationCount, expectedPopulationCount)
	}
}
