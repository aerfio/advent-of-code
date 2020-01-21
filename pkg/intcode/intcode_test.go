package intcode

import (
	"fmt"
	"testing"

	"github.com/aerfio/advent-of-code/pkg/input/task2"
	"github.com/onsi/gomega"
)

func Test_handleOperation(t *testing.T) {
	type args struct {
		position int
		arg      []int
	}
	tests := []struct {
		args         args
		wantSum      *[]int
		wantMultiply *[]int
	}{
		{

			args: args{
				position: 0,
				arg: []int{
					1, 0, 0, 0,
				},
			},
			wantSum:      &[]int{2, 0, 0, 0},
			wantMultiply: &[]int{1, 0, 0, 0},
		},
		{

			args: args{
				position: 0,
				arg: []int{
					2, 3, 0, 3,
				},
			},
			wantSum:      &[]int{2, 3, 0, 5},
			wantMultiply: &[]int{2, 3, 0, 6},
		},
		{

			args: args{
				position: 0,
				arg: []int{
					1, 1, 1, 4, 99, 5, 6, 0, 99},
			},
			wantSum: &[]int{
				1, 1, 1, 4, 2, 5, 6, 0, 99,
			},
			wantMultiply: &[]int{
				1, 1, 1, 4, 1, 5, 6, 0, 99,
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v\n", tt.args.arg), func(t *testing.T) {
			gm := gomega.NewGomegaWithT(t)
			gm.Expect(handleOperation(tt.args.position, tt.args.arg, Sum)).To(gomega.Equal(tt.wantSum))
			gm.Expect(handleOperation(tt.args.position, tt.args.arg, Multiply)).To(gomega.Equal(tt.wantMultiply))
		})
	}
}

func Test_program_Run(t *testing.T) {
	type testData struct {
		in      string
		out     []int
		restore *RunOpts
	}

	tests := []struct {
		data    testData
		wantErr bool
	}{
		{
			data: testData{
				in: `1,0,0,0,99`,
				out: []int{
					2, 0, 0, 0, 99,
				},
				restore: nil,
			},
			wantErr: false,
		},
		{
			data: testData{
				in:      `2,3,0,3,99`,
				out:     []int{2, 3, 0, 6, 99},
				restore: nil,
			},

			wantErr: false,
		},
		{
			data: testData{
				in: `2,4,4,5,99,0`,
				out: []int{
					2, 4, 4, 5, 99, 9801,
				},
				restore: nil,
			},
			wantErr: false,
		},
		{
			data: testData{
				in: `1,1,1,4,99,5,6,0,99`,
				out: []int{
					30, 1, 1, 4, 2, 5, 6, 0, 99,
				},
				restore: nil,
			},
			wantErr: false,
		},
		{
			data: testData{
				in: `1,9,10,3,2,3,11,0,99,30,40,50`,
				out: []int{
					3500, 9, 10, 70,
					2, 3, 11, 0,
					99, 30, 40, 50,
				},
				restore: nil,
			},
			wantErr: false,
		},
		{
			data: testData{
				in:      task2.Data,
				out:     myCaseAnswer,
				restore: &RunOpts{InitialVerb: 12, InitialNoun: 2},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.data.in), func(t *testing.T) {
			gm := gomega.NewGomegaWithT(t)
			prog, err := New(tt.data.in)
			if tt.wantErr {
				gm.Expect(err).To(gomega.HaveOccurred())
			} else {
				gm.Expect(err).NotTo(gomega.HaveOccurred())
			}

			prog.Run(tt.data.restore)

			gm.Expect(prog.parsed).To(gomega.Equal(tt.data.out))
		})
	}
}

var myCaseAnswer = []int{3760627, 12, 2, 2, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 9, 36, 1, 10, 19, 40, 2, 9, 23, 120, 1, 6, 27, 122, 2, 31, 9, 366, 1, 5, 35, 367, 1, 10, 39, 371, 1, 10, 43, 375, 2, 13, 47, 1875, 1, 10, 51, 1879, 2, 55, 10, 7516, 1, 9, 59, 7519, 2, 6, 63, 15038, 1, 5, 67, 15039, 1, 71, 5, 15040, 1, 5, 75, 15041, 2, 79, 13, 75205, 1, 83, 5, 75206, 2, 6, 87, 150412, 1, 5, 91, 150413, 1, 95, 9, 150416, 1, 99, 6, 150418, 1, 103, 13, 150423, 1, 107, 5, 150424, 2, 111, 13, 752120, 1, 115, 6, 752122, 1, 6, 119, 752124, 2, 123, 13, 3760620, 1, 10, 127, 3760624, 1, 131, 2, 3760626, 1, 135, 5, 0, 99, 2, 14, 0, 0}
