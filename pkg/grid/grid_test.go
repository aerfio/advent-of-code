package grid

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func Test_stringToDir(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		want direction
	}{
		{
			name: "correctly returns up given u",
			want: up,
			args: args{token: "U"},
		},

		{
			name: "correctly returns down given d",
			want: down,
			args: args{token: "D"},
		},

		{
			name: "correctly returns right given r",
			want: right,
			args: args{token: "R"},
		},

		{
			name: "correctly returns left given L",
			want: left,
			args: args{token: "L"},
		},

		{
			name: "correctly returns up given random letters",
			want: up,
			args: args{token: "xD"},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s - lowercased", tt.name), func(t *testing.T) {
			if got := stringToDir(strings.ToLower(tt.args.token)); got != tt.want {
				t.Errorf("stringToDir() = %v, want %v", got, tt.want)
			}
		})
		t.Run(fmt.Sprintf("%s - uppercased", tt.name), func(t *testing.T) {
			if got := stringToDir(strings.ToUpper(tt.args.token)); got != tt.want {
				t.Errorf("stringToDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		arg string
	}
	tests := []struct {
		name    string
		args    args
		want    []movement
		wantErr bool
	}{
		{
			name:    "R997",
			want:    []movement{{distance: 997, direction: right}},
			wantErr: false,
			args:    args{arg: "R997"},
		},
		{
			name:    "simple",
			want:    []movement{},
			wantErr: true,
			args:    args{arg: "RXD"},
		},
		{
			name: "U20,R12,D36",
			want: []movement{
				{distance: 20, direction: up}, {
					direction: right,
					distance:  12,
				}, {
					direction: down,
					distance:  36,
				}},
			wantErr: false,
			args:    args{arg: "U20,R12,D36"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_movementSlice_toGrid(t *testing.T) {
	tests := []struct {
		name string
		ms   movementSlice
		want []point
	}{
		{
			name: "blah",
			ms:   movementSlice{{distance: 2, direction: up}, {direction: right, distance: 3}},
			want: []point{{x: 0, y: 0}, {x: 0, y: 2}, {x: 3, y: 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ms.toCornerPoints(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toCornerPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
