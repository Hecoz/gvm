package classfile

type ConstantMemberrefInfo struct {

	cp					ConstantPool
	classIndex			uint16
	nameAndTypeIndex	uint16
}


//go 语言中没有继承，可以通过结构体嵌套来模拟
type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {

	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberrefInfo) NameAndDescriptro() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

