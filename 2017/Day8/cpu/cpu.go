package cpu

import (
	"strconv"
)

type registers struct {
	registers map[string]int
	largest   int
}

var r registers

func init() {
	r.registers = make(map[string]int)
	r.largest = -1
}

type Instruction struct {
	Register         string
	ModifyOperation  string
	ModifyValue      int
	CompareA         string
	CompareOperation string
	CompareB         int
	modify           bool
}

func NewInstruction(instruction []string) (*Instruction, error) {
	return &Instruction{
		Register:         instruction[0],
		ModifyOperation:  instruction[1],
		ModifyValue:      toInt(instruction[2]),
		CompareA:         instruction[4],
		CompareOperation: instruction[5],
		CompareB:         toInt(instruction[6]),
	}, nil
}

func (i *Instruction) ProcessInstruction() {
	switch i.CompareOperation {
	case "<":
		i.lessThan()
	case "<=":
		i.lessThanEqual()
	case ">":
		i.greaterThan()
	case ">=":
		i.greaterThanEqual()
	case "==":
		i.equal()
	case "!=":
		i.notEqual()
	}
	if i.modify {
		if i.ModifyOperation == "inc" {
			r.registers[i.Register] += i.ModifyValue
		} else {
			r.registers[i.Register] -= i.ModifyValue
		}
	}
	if r.registers[i.Register] > r.largest {
		r.largest = r.registers[i.Register]
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func LargestRegister() (string, int) {
	largest := -1
	var register string
	for k, v := range r.registers {
		if v > largest {
			largest = v
			register = k
		}
	}
	return register, largest
}

func Highmark() int {
	return r.largest
}

func (i *Instruction) lessThan() {
	i.modify = r.registers[i.CompareA] < i.CompareB
}

func (i *Instruction) lessThanEqual() {
	i.modify = r.registers[i.CompareA] <= i.CompareB
}

func (i *Instruction) greaterThan() {
	i.modify = r.registers[i.CompareA] > i.CompareB
}

func (i *Instruction) greaterThanEqual() {
	i.modify = r.registers[i.CompareA] >= i.CompareB
}

func (i *Instruction) equal() {
	i.modify = r.registers[i.CompareA] == i.CompareB
}

func (i *Instruction) notEqual() {
	i.modify = r.registers[i.CompareA] != i.CompareB
}
