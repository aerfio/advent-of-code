package fuel

import (
	"testing"
)

func Test_compute(t *testing.T) {
	type args struct {
		arg int
		acc int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return 2 for 14",
			want: 2,
			args: args{
				arg: 14,
				acc: 0,
			},
		},
		{
			name: "should return 966 for 1969",
			want: 966,
			args: args{
				arg: 1969,
				acc: 0,
			},
		}, {
			name: "should return 50346 for 100756",
			want: 50346,
			args: args{
				arg: 100756,
				acc: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeRec(tt.args.arg, tt.args.acc); got != tt.want {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compute1(t *testing.T) {

	tests := []struct {
		arg  int
		want int
	}{
		{
			want: 2,
			arg:  14,
		},
		{
			arg:  1969,
			want: 654,
		},
		{
			arg:  654,
			want: 216,
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.arg), func(t *testing.T) {
			if got := compute(tt.arg); got != tt.want {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}
