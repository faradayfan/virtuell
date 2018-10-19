package main

import (
	"testing"
)

func TestRun_Add_1and1(t *testing.T) {
	conf := Config{
		CodeMemSize: 8,
		DataMemSize: 8,
		StackSize:   8,
	}
	spec := Create(conf)

	spec.LoadCode([]int{ICONST, 1, ICONST, 1, 1, HALT})

	spec.Run()

	if spec.ReadStack() != 2 {
		t.Fail()
	}
}

func TestRun_Sub_3from2(t *testing.T) {
	conf := Config{
		CodeMemSize: 8,
		DataMemSize: 8,
		StackSize:   8,
	}
	spec := Create(conf)

	spec.LoadCode([]int{ICONST, 2, ICONST, 3, 2, HALT})

	spec.Run()

	if spec.ReadStack() != -1 {
		t.Fail()
	}
}

func TestRun_Mul_2and3(t *testing.T) {
	conf := Config{
		CodeMemSize: 8,
		DataMemSize: 8,
		StackSize:   8,
	}
	spec := Create(conf)

	spec.LoadCode([]int{ICONST, 2, ICONST, 3, 3, HALT})

	spec.Run()

	if spec.ReadStack() != 6 {
		t.Fail()
	}
}

func TestRun_Div_8by4(t *testing.T) {
	conf := Config{
		CodeMemSize: 8,
		DataMemSize: 8,
		StackSize:   8,
	}
	spec := Create(conf)

	spec.LoadCode([]int{ICONST, 8, ICONST, 4, I_DIV, HALT})

	spec.Run()

	if spec.ReadStack() != 2 {
		t.Fail()
	}
}

// func TestRun_ILT(t *testing.T) {
// 	conf := Config{
// 		CodeMemSize: 8,
// 		DataMemSize: 8,
// 		StackSize:   8,
// 	}
// 	spec := Create(conf)

// 	spec.LoadCode([]int{ICONST, 3, 5, 4, 19})

// 	spec.Run()

// 	if spec.ReadInstructionPointer() != 5 {
// 		t.Fail()
// 	}
// }

func TestRun_BR(t *testing.T) {
	conf := Config{
		CodeMemSize: 8,
		DataMemSize: 8,
		StackSize:   8,
	}
	spec := Create(conf)

	spec.LoadCode([]int{BR, 3, 0, HALT})

	spec.Run()

	if spec.ReadInstructionPointer() != 3 {
		t.Fail()
	}
}

func TestRun_BRT(t *testing.T) {
	conf := Config{
		CodeMemSize: 8,
		DataMemSize: 8,
		StackSize:   8,
	}
	spec := Create(conf)

	spec.LoadCode([]int{ICONST, 1, BRT, 5, HALT, HALT})

	spec.Run()

	if spec.ReadInstructionPointer() != 5 {
		t.Fail()
	}
}

func TestRun_BRF(t *testing.T) {
	conf := Config{
		CodeMemSize: 8,
		DataMemSize: 8,
		StackSize:   8,
	}
	spec := Create(conf)

	spec.LoadCode([]int{ICONST, 0, BRF, 5, HALT, HALT})

	spec.Run()

	if spec.ReadInstructionPointer() != 5 {
		t.Fail()
	}
	if spec.ReadStack() != 0 {
		t.Fail()
	}
}
