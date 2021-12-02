package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_readCommands(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []Command
		wantErr bool
	}{
		{
			name: "simple case",
			args: args{
				input: strings.NewReader("forward 1\n"),
			},
			want: []Command{
				{
					Direction: Forward,
					Distance:  1,
				},
			},
			wantErr: false,
		},
		{
			name: "website example",
			args: args{
				input: strings.NewReader("forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2\n"),
			},
			want: []Command{
				{
					Direction: Forward,
					Distance:  5,
				},
				{
					Direction: Down,
					Distance:  5,
				},
				{
					Direction: Forward,
					Distance:  8,
				},
				{
					Direction: Up,
					Distance:  3,
				},
				{
					Direction: Down,
					Distance:  8,
				},
				{
					Direction: Forward,
					Distance:  2,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readCommands(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("readCommands() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readCommands() got = %v, want %v", got, tt.want)
			}
		})
	}
}
