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
