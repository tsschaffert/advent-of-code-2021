package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_calculateGammaRate(t *testing.T) {
	type args struct {
		report DiagnosticReport
	}
	tests := []struct {
		name string
		args args
		want ReportRow
	}{
		{
			name: "website example",
			args: args{
				report: []ReportRow{
					{false, false, true, false, false},
					{true, true, true, true, false},
					{true, false, true, true, false},
					{true, false, true, true, true},
					{true, false, true, false, true},
					{false, true, true, true, true},
					{false, false, true, true, true},
					{true, true, true, false, false},
					{true, false, false, false, false},
					{true, true, false, false, true},
					{false, false, false, true, false},
					{false, true, false, true, false},
				},
			},
			want: ReportRow{true, false, true, true, false},
		},
		{
			name: "simple example",
			args: args{
				report: []ReportRow{
					{false, true},
					{false, true},
				},
			},
			want: ReportRow{false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateGammaRate(tt.args.report); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateGammaRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToGammaRate(t *testing.T) {
	type args struct {
		row ReportRow
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0 test",
			args: args{
				[]bool{false},
			},
			want: 0,
		},
		{
			name: "1 test",
			args: args{
				[]bool{true},
			},
			want: 1,
		},
		{
			name: "simple test",
			args: args{
				[]bool{true, false},
			},
			want: 2,
		},
		{
			name: "website test",
			args: args{
				[]bool{true, false, true, true, false},
			},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToNumber(tt.args.row); got != tt.want {
				t.Errorf("convertToGammaRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readDiagnosticReport(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    DiagnosticReport
		wantErr bool
	}{
		{
			name: "simple test",
			args: args{
				input: strings.NewReader("01\n01"),
			},
			want: DiagnosticReport{
				ReportRow{false, true},
				ReportRow{false, true},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readDiagnosticReport(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("readDiagnosticReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readDiagnosticReport() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDeltaRate(t *testing.T) {
	type args struct {
		gammaRate ReportRow
	}
	tests := []struct {
		name string
		args args
		want ReportRow
	}{
		{
			name: "simple test",
			args: args{
				gammaRate: ReportRow{false, true},
			},
			want: ReportRow{true, false},
		},
		{
			name: "website test",
			args: args{
				gammaRate: ReportRow{true, false, true, true, false},
			},
			want: ReportRow{false, true, false, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDeltaRate(tt.args.gammaRate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDeltaRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
