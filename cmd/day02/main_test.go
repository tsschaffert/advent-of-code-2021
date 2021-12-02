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

func TestPosition_ApplyCommand(t *testing.T) {
	type fields struct {
		Horizontal int
		Depth      int
	}
	type args struct {
		command Command
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Position
	}{
		{
			name: "move up",
			fields: fields{
				Horizontal: 0,
				Depth:      10,
			},
			args: args{
				Command{
					Direction: Up,
					Distance:  1,
				},
			},
			want: Position{
				Horizontal: 0,
				Depth:      9,
			},
		},
		{
			name: "move down",
			fields: fields{
				Horizontal: 0,
				Depth:      10,
			},
			args: args{
				Command{
					Direction: Down,
					Distance:  1,
				},
			},
			want: Position{
				Horizontal: 0,
				Depth:      11,
			},
		},
		{
			name: "move forward",
			fields: fields{
				Horizontal: 0,
				Depth:      10,
			},
			args: args{
				Command{
					Direction: Forward,
					Distance:  1,
				},
			},
			want: Position{
				Horizontal: 1,
				Depth:      10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Position{
				Horizontal: tt.fields.Horizontal,
				Depth:      tt.fields.Depth,
			}
			if got := p.ApplyCommand(tt.args.command); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApplyCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_applyCommands(t *testing.T) {
	type args struct {
		initialPosition Position
		commands        []Command
	}
	tests := []struct {
		name string
		args args
		want Position
	}{
		{
			name: "website example",
			args: args{
				initialPosition: Position{0, 0},
				commands: []Command{
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
			},
			want: Position{15, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyCommands(tt.args.initialPosition, tt.args.commands); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyCommands() = %v, want %v", got, tt.want)
			}
		})
	}
}
