package grid

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/onsi/gomega"
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
			gm := gomega.NewGomegaWithT(t)

			got, err := parse(tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gm.Expect(got).To(gomega.BeEquivalentTo(tt.want))
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
			gm := gomega.NewGomegaWithT(t)
			gm.Expect(tt.ms.toCornerPoints()).To(gomega.BeEquivalentTo(tt.want))
		})
	}
}

func Test_getPathBetweenPoints(t *testing.T) {
	type args struct {
		p1 point
		p2 point
	}
	tests := []struct {
		name    string
		args    args
		want    path
		wantErr bool
	}{
		{
			name:    "simple - OX",
			wantErr: false,
			args: args{
				p1: point{x: 0, y: 0},
				p2: point{x: 4, y: 0},
			},
			want: path{{x: 1, y: 0}, {x: 2, y: 0}, {x: 3, y: 0}},
		},
		{
			name:    "simple,inverted - OX",
			wantErr: false,
			args: args{
				p1: point{x: 4, y: 0},
				p2: point{x: 0, y: 0},
			},
			want: path{{x: 3, y: 0}, {x: 2, y: 0}, {x: 1, y: 0}},
		},
		{
			name:    "simple - OY",
			wantErr: false,
			args: args{
				p1: point{x: 0, y: 0},
				p2: point{x: 0, y: 4},
			},
			want: path{{x: 0, y: 1}, {x: 0, y: 2}, {x: 0, y: 3}},
		},
		{
			name:    "simple,inverted - OY",
			wantErr: false,
			args: args{
				p1: point{x: 0, y: 4},
				p2: point{x: 0, y: 0},
			},
			want: path{{x: 0, y: 3}, {x: 0, y: 2}, {x: 0, y: 1}},
		},
		{
			name:    "only numbers below zero - OX",
			wantErr: false,
			args: args{
				p1: point{x: -10, y: -10},
				p2: point{x: -14, y: -10},
			},
			want: path{{x: -11, y: -10}, {x: -12, y: -10}, {x: -13, y: -10}},
		},
		{
			name:    "only numbers below zero - OY",
			wantErr: false,
			args: args{
				p1: point{x: -10, y: -10},
				p2: point{x: -10, y: -14},
			},
			want: path{{x: -10, y: -11}, {x: -10, y: -12}, {x: -10, y: -13}},
		},
		{
			name:    "retuns error",
			wantErr: true,
			args: args{
				p1: point{x: 0, y: 0},
				p2: point{x: 0, y: 0},
			},
			want: path{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPathBetweenPoints(tt.args.p1, tt.args.p2)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPathBetweenPoints() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPathBetweenPoints() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_point_manhattanDistance(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	type args struct {
		next point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "second point positive",
			want: 10,
			args: args{next: point{x: 5, y: 5}},
			fields: fields{
				x: 0,
				y: 0,
			},
		},
		{
			name: "second point negative",
			want: 10,
			args: args{next: point{x: -5, y: -5}},
			fields: fields{
				x: 0,
				y: 0,
			},
		},
		{
			name: "all negative",
			want: 10,
			args: args{next: point{x: -10, y: -10}},
			fields: fields{
				x: -15,
				y: -15,
			},
		},
		{
			name: "all negative - reverse order",
			want: 10,
			args: args{next: point{x: -15, y: -15}},
			fields: fields{
				x: -10,
				y: -10,
			},
		},
		{
			name: "all positive",
			want: 12,
			args: args{next: point{x: 1, y: 1}},
			fields: fields{
				x: 7,
				y: 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := point{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := p.manhattanDistance(tt.args.next); got != tt.want {
				t.Errorf("manhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_point_isBetweenTwoPoints(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	type args struct {
		p1 point
		p2 point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "is between two points on OX axis",
			want: true,
			fields: fields{
				x: 0,
				y: 0,
			},
			args: args{
				p1: point{x: 5, y: 0},
				p2: point{x: -5, y: 0},
			},
		},
		{
			name: "is between two points on OX axis - reversed order",
			want: true,
			fields: fields{
				x: 0,
				y: 0,
			},
			args: args{
				p1: point{x: -5, y: 0},
				p2: point{x: 5, y: 0},
			},
		},
		{
			name: "is between two points on OY axis",
			want: true,
			fields: fields{
				x: 0,
				y: 0,
			},
			args: args{
				p1: point{x: 0, y: -5},
				p2: point{x: 0, y: 5},
			},
		},
		{
			name: "is between two points on OY axis - reversed order",
			want: true,
			fields: fields{
				x: 0,
				y: 0,
			},
			args: args{
				p1: point{x: 0, y: 5},
				p2: point{x: 0, y: -5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := point{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := p.isBetweenTwoPoints(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("isBetweenTwoPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
