package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_readSignals(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []Reading
		wantErr bool
	}{
		{
			name: "website example",
			args: args{
				input: strings.NewReader("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf\n"),
			},
			want: []Reading{
				{
					signalPatterns: []Signal{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
					output:         []Signal{"cdfeb", "fcadb", "cdfeb", "cdbaf"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readSignals(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("readSignals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readSignals() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countUniqueSignals(t *testing.T) {
	type args struct {
		readings []Reading
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "shortened website example",
			args: args{
				readings: []Reading{
					{
						signalPatterns: []Signal{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
						output:         []Signal{"fdgacbe", "cefdb", "cefbgd", "gcbe"},
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countUniqueSignals(tt.args.readings); got != tt.want {
				t.Errorf("countUniqueSignals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_analyseMappings(t *testing.T) {
	type args struct {
		readings []Reading
	}
	tests := []struct {
		name string
		args args
		want []Reading
	}{
		{
			name: "website example",
			args: args{
				[]Reading{
					{
						signalPatterns: []Signal{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
						output:         []Signal{"cdfeb", "fcadb", "cdfeb", "cdbaf"},
						mapping:        nil,
					},
				},
			},
			want: []Reading{
				{
					signalPatterns: []Signal{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
					output:         []Signal{"cdfeb", "fcadb", "cdfeb", "cdbaf"},
					mapping: map[Signal]int{
						"acedgfb": 8,
						"cdfbe":   5,
						"gcdfa":   2,
						"fbcad":   3,
						"dab":     7,
						"cefabd":  9,
						"cdfgeb":  6,
						"eafb":    4,
						"cagedb":  0,
						"ab":      1,
					},
				},
			},
		},
		{
			name: "3 and 5 overlap example",
			args: args{
				[]Reading{
					{
						signalPatterns: []Signal{"aceg", "gdaef", "fcbegda", "dcefab", "afcedg", "abgdfc", "cadef", "agf", "degbf", "ag"},
						output:         []Signal{"bdgfe", "bacfgd", "cadfe", "fcgeabd"},
						mapping:        nil,
					},
				},
			},
			want: []Reading{
				{
					signalPatterns: []Signal{"aceg", "gdaef", "fcbegda", "dcefab", "afcedg", "abgdfc", "cadef", "agf", "degbf", "ag"},
					output:         []Signal{"bdgfe", "bacfgd", "cadfe", "fcgeabd"},
					mapping: map[Signal]int{
						"fcbegda": 8,
						"degbf":   2,
						"cadef":   5,
						"gdaef":   3,
						"agf":     7,
						"afcedg":  9,
						"dcefab":  6,
						"aceg":    4,
						"abgdfc":  0,
						"ag":      1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := analyseMappings(tt.args.readings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("analyseMappings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignal_equals(t *testing.T) {
	type args struct {
		other Signal
	}
	tests := []struct {
		name string
		s    Signal
		args args
		want bool
	}{
		{
			name: "simple match",
			s:    "a",
			args: args{other: "a"},
			want: true,
		},
		{
			name: "simple non-match",
			s:    "a",
			args: args{other: "b"},
			want: false,
		},
		{
			name: "transposition match",
			s:    "ab",
			args: args{other: "ba"},
			want: true,
		},
		{
			name: "transposition non-match",
			s:    "ab",
			args: args{other: "ca"},
			want: false,
		},
		{
			name: "different length",
			s:    "ab",
			args: args{other: "abc"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.equals(tt.args.other); got != tt.want {
				t.Errorf("equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignal_countOverlap(t *testing.T) {
	type args struct {
		other Signal
	}
	tests := []struct {
		name string
		s    Signal
		args args
		want int
	}{
		{
			name: "simple test",
			s:    "a",
			args: args{other: "a"},
			want: 1,
		},
		{
			name: "simple test",
			s:    "a",
			args: args{other: "b"},
			want: 0,
		},
		{
			name: "transposition match",
			s:    "ab",
			args: args{other: "ba"},
			want: 2,
		},
		{
			name: "transposition non-match",
			s:    "ab",
			args: args{other: "ca"},
			want: 1,
		},
		{
			name: "different length",
			s:    "ab",
			args: args{other: "abc"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.countOverlap(tt.args.other); got != tt.want {
				t.Errorf("countOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMatch(t *testing.T) {
	type args struct {
		signalPatterns []Signal
		reverseMapping map[int]Signal
		matches        func(Signal, map[int]Signal) bool
	}
	tests := []struct {
		name string
		args args
		want Signal
	}{
		{
			name: "simple test",
			args: args{
				signalPatterns: []Signal{"a"},
				reverseMapping: nil,
				matches: func(signal Signal, m map[int]Signal) bool {
					return signal == "a"
				},
			},
			want: "a",
		},
		{
			name: "website example",
			args: args{
				signalPatterns: []Signal{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
				reverseMapping: nil,
				matches: func(signal Signal, m map[int]Signal) bool {
					return len(signal) == 2
				},
			},
			want: "ab",
		},
		{
			name: "website example with existing mapping",
			args: args{
				signalPatterns: []Signal{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
				reverseMapping: map[int]Signal{7: "dab"},
				matches: func(signal Signal, m map[int]Signal) bool {
					return len(signal) == 5 && signal.countOverlap(m[7]) == len(m[7])
				},
			},
			want: "fbcad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMatch(tt.args.signalPatterns, tt.args.reverseMapping, tt.args.matches); got != tt.want {
				t.Errorf("findMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToNumber(t *testing.T) {
	type args struct {
		reading Reading
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "website example",
			args: args{
				reading: Reading{
					signalPatterns: []Signal{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
					output:         []Signal{"cdfeb", "fcadb", "cdfeb", "cdbaf"},
					mapping: map[Signal]int{
						"acedgfb": 8,
						"cdfbe":   5,
						"gcdfa":   2,
						"fbcad":   3,
						"dab":     7,
						"cefabd":  9,
						"cdfgeb":  6,
						"eafb":    4,
						"cagedb":  0,
						"ab":      1,
					},
				},
			},
			want: 5353,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToNumber(tt.args.reading); got != tt.want {
				t.Errorf("convertToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lookupInMap(t *testing.T) {
	type args struct {
		output  Signal
		mapping map[Signal]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple test",
			args: args{
				output:  "a",
				mapping: map[Signal]int{"a": 1},
			},
			want: 1,
		},
		{
			name: "another test",
			args: args{
				output:  "ab",
				mapping: map[Signal]int{"a": 1, "ba": 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lookupInMap(tt.args.output, tt.args.mapping); got != tt.want {
				t.Errorf("lookupInMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
