package main

import (
	"fmt"
)

// iadd - integer add 2 operands on stack
// isub - integer sub 2 operands on stack
// imul - integer multiply 2 operands on stack
// idiv - integer multiply 2 operands on stack
// ilt - integer less than
// ieq - integer equal
// br addr - jump to address
// brt addr - branch if true
// brf addr - branch if false
// iconst value - push integer const
// load addr - load local(code)
// gload addr - load global(data)
// store addr - store local(code)
// gstore addr - store global(data)
// print -
// pop - toss out top of stack
// call addr, numArgs - call addr, expect numArgs
// ret - return from function
// halt
const I_ADD = 1
const I_SUB = 2
const I_MUL = 3
const I_DIV = 4
const ILT = 5
const IEQ = 6
const BR = 7
const BRT = 8
const BRF = 9
const ICONST = 10
const LOAD = 11
const GLOAD = 12
const STORE = 13
const GSTORE = 14
const PRINT = 15
const POP = 16
const CALL = 17
const RET = 18
const HALT = 19

type VM interface {
	Run()
	LoadCode(code []int)
	LoadData(data []int)
	ReadStack() int
	ReadInstructionPointer() int
	ReadStackPointer() int
	ReadFramePointer() int
}

type Config struct {
	CodeMemSize int
	DataMemSize int
	StackSize   int
}

type mem struct {
	config             Config
	stack              []int
	codemem            []int
	datamem            []int
	stackpointer       int
	instructionpointer int
	framepointer       int
}

func (m *mem) Run() {
	bytecode := m.codemem[m.instructionpointer]
	for bytecode != HALT && m.instructionpointer < m.config.CodeMemSize {
		m.instructionpointer += 1
		m.decode(bytecode)
		bytecode = m.codemem[m.instructionpointer]
	}
}

func (m *mem) LoadCode(code []int) {
	copy(m.codemem, code)
}

func (m *mem) LoadData(data []int) {
	copy(m.datamem, data)
}

func (m *mem) ReadStack() int {
	return m.stack[m.stackpointer]
}

func (m *mem) ReadInstructionPointer() int {
	return m.instructionpointer
}

func (m *mem) ReadStackPointer() int {
	return m.stackpointer
}

func (m *mem) ReadFramePointer() int {
	return m.framepointer
}

func (m *mem) decode(instruction int) {
	switch instruction {
	case I_ADD:
		m.stack[m.stackpointer-1] = m.stack[m.stackpointer-1] + m.stack[m.stackpointer]
		m.stackpointer -= 1
		break
	case I_SUB:
		m.stack[m.stackpointer-1] = m.stack[m.stackpointer-1] - m.stack[m.stackpointer]
		m.stackpointer -= 1
		break
	case I_MUL:
		m.stack[m.stackpointer-1] = m.stack[m.stackpointer-1] * m.stack[m.stackpointer]
		m.stackpointer -= 1
		break
	case I_DIV:
		m.stack[m.stackpointer-1] = m.stack[m.stackpointer-1] / m.stack[m.stackpointer]
		m.stackpointer -= 1
		break
	case ILT:
		break
	case IEQ:
		break
	case BR:
		m.instructionpointer = m.codemem[m.instructionpointer]
		break
	case BRT:
		if m.stack[m.stackpointer] != 0 {
			m.instructionpointer = m.codemem[m.instructionpointer]
		} else {
			m.instructionpointer += 1
		}
		break
	case BRF:
		if m.stack[m.stackpointer] == 0 {
			m.instructionpointer = m.codemem[m.instructionpointer]
		} else {
			m.instructionpointer += 1
		}
		break
	case ICONST:
		m.stackpointer += 1
		m.stack[m.stackpointer] = m.codemem[m.instructionpointer]
		m.instructionpointer += 1
		break
	case LOAD:
		break
	case GLOAD:
		break
	case STORE:
		break
	case GSTORE:
		break
	case PRINT:
		fmt.Printf("%d", m.stack[m.stackpointer])
		break
	case POP:
		break
	case CALL:
		break
	case RET:
		break
	case HALT:
		break
	default:
		panic(fmt.Sprintf("Could not interpret bytecode '%d' as instruction at ip '%d'", instruction, m.instructionpointer))
	}
}

func Create(c Config) VM {
	return &mem{
		config:             c,
		stack:              make([]int, c.StackSize),
		datamem:            make([]int, c.DataMemSize),
		codemem:            make([]int, c.CodeMemSize),
		stackpointer:       -1,
		instructionpointer: 0,
		framepointer:       0,
	}
}
