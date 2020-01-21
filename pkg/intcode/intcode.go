package intcode

import (
	"strconv"
	"strings"
)

type program struct {
	parsed  []int
	initial []int
}

func New(code string) (*program, error) {
	data, err := parse(code)
	if err != nil {
		return nil, err
	}
	return &program{parsed: data, initial: data}, nil
}

func parse(code string) ([]int, error) {
	parsed := make([]int, 0)
	for _, num := range strings.Split(code, ",") {
		data, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, data)
	}
	return parsed, nil
}

func (p *program) restoreGravityAssistProgram(opts RunOpts) {
	(*p).parsed[1] = opts.InitialNoun
	(*p).parsed[2] = opts.InitialVerb
}

type RunOpts struct {
	InitialVerb int
	InitialNoun int
}

func (p *program) Run(opts *RunOpts) {
	if opts != nil {
		p.restoreGravityAssistProgram(*opts)
	}

	// currPos := 0
	numberOfSubPrograms := len(p.parsed) / 4
	for i := 0; i < numberOfSubPrograms; i++ {
		position := i * 4
		currOpcode := opcode(p.parsed[position])

		if currOpcode == Break {
			break
		}
		(*p).parsed = *handleOperation(position, p.parsed, currOpcode)
	}
}

func (p *program) GetOutput() int {
	return p.parsed[0]
}

func (p *program) Reset() {
	(*p).parsed = p.initial
}

type opcode int

const (
	Sum      opcode = 1
	Multiply opcode = 2
	Break    opcode = 99
)

func handleOperation(position int, arg []int, oper opcode) *[]int {
	handle := make([]int, len(arg))
	copy(handle, arg)

	num1 := handle[position+1]
	num2 := handle[position+2]
	target := handle[position+3]

	switch oper {
	case Multiply:
		handle[target] = handle[num1] * handle[num2]
	case Sum:
		handle[target] = handle[num1] + handle[num2]
	default:
		handle[target] = handle[num1] + handle[num2]
	}

	return &handle
}
