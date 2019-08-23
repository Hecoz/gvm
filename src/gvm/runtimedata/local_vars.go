package runtimedata

import "math"

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {

	if maxLocals > 0 {

		return make([]Slot, maxLocals)
	}
	return nil
}

func (self LocalVars) SetInt(index uint, val int32)  {
	self[index].num = val
}

func (self LocalVars) GetInt(index uint) int32{

	return self[index].num
}

func (self LocalVars) SetFloat(index uint, val float32) {

	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}