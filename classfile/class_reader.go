package classfile

import (
	"encoding/binary"
)

type ClassReader struct {

	data []byte
}

func (self *ClassReader) readUint8() uint8 {

	//每次取0索引位置变量，一个byte刚好是8位
	val := self.data[0]
	//重新赋值data，有个疑问？这样的话如果取到一半不取了，剩下的数据岂不是不完整了？
	self.data = self.data[1:]
	return val
}

func (self *ClassReader) readUint16() uint16 {

	//Go标准库 encoding/binary中定义的变量BigEndian可以调用不同的参数从[]byte中解码出不同字节大小的数据
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

func (self *ClassReader) readUint32() uint32 {

	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readUint64() uint64 {

	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

/**
读取uint16表，表的大小由开头的uint16数据指出
*/
func (self *ClassReader) readUint16s() []uint16 {

	tableSize := self.readUint16()
	table := make([]uint16,tableSize)
	for i := range table{
		table[i] = self.readUint16()
	}
	return table
}

/*
读取指定大小的字节
*/
func (self *ClassReader) readBytes(byteSize uint32) []byte {

	bytes := self.data[:byteSize]
	self.data = self.data[byteSize:]
	return bytes
}