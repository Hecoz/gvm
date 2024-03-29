package classfile

/*
java常量数据中常量的格式不尽相同，常量数据的第一个字节是tag用来区分不同的常量。
JAVA虚拟机规范给出的常量结构：
cp_info{
	u1	tag;
	u1	info[];
}
 */
//tag 常量值定义
const(
	CONSTANT_Utf8					= 1
	CONSTANT_Integer				= 3
	CONSTANT_Float					= 4
	CONSTANT_Long					= 5
	CONSTANT_Double					= 6
	CONSTANT_Class					= 7
	CONSTANT_String					= 8
	CONSTANT_Fieldref    			= 9
	CONSTANT_Methodref				= 10
	CONSTANT_InterfaceMethodref		= 11
	CONSTANT_NameAndType			= 12
	CONSTANT_MethodHandle			= 15
	CONSTANT_MethodType				= 16
	CONSTANT_InvokeDynamic			= 18
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {

	//读取一个字节的tag,用来判断是什么数据
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {

	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{}
	case CONSTANT_Class:
		return &ConstantClassInfo{}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}