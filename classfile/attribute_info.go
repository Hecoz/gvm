package classfile


/*
attribute_info{
	u2	attribute_name_index;
	u4	attribute_length;
	u1	info[attribute_length];
}
 */
type AttributeInfo interface {

	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {

	// 01 读取属性表大小
	attributesCount := reader.readUint16()
	// 02 读取属性信息
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes{

		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {

	// 01 读取属性名索引，总共有23中属性信息， 目前知道的属性名有两种，一种是在方法中叫CODE，一种是在类中叫SourceFile
	attrNameIndex := reader.readUint16()
	//    从常量池中读取属性名
	attrName := cp.getUtf8(attrNameIndex)
	// 02 读取属性表长度  4字节
	attrLen := reader.readUint32()
	// 03 读取属性表信息，属性表分为，以下9个
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {

	switch attrName {

	case "Code":				return &CodeAttribute{cp: cp}
	case "ConstantValue":		return &ConstantValueAttribute{}
	case "Deprecated":			return &DeprecatedAttribute{}
	case "Exceptions":			return &ExceptionsAttribute{}
	case "LineNumberTable":		return &LineNumberTableAttribute{}
	case "LocalVariableTable":	return &LocalVariableTableAttribute{}
	case "SourceFile":			return &SourceFileAttribute{cp: cp}
	case "Synthetic":			return &SyntheticAttribute{}
	default:					return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
