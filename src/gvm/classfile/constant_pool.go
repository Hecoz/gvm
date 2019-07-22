package classfile

type ConstantPool []ConstantInfo

/*
读取常量池
注意：
（1）常量池的索引从1开始，到 n - 1结束，所以是[1, n)
（2）CONSTANT_Long_info, CONSTANT_Double_info占两个位置
 */
func readConstantPool(reader *ClassReader) ConstrantPool {

	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)

	//inde begin with 1
	for i := 1; i < cpCount; i++{

		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {

		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

func (self ConstantPool) getConstantInfo(index uint16) ConstrantPool {

	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {

	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (self ConstantPool) getClassName(index uint16) string{

	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

func (self ConstantPool) getUtf8(index uint16) string {

	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	reutrn utf8Info.str
}