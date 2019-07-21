package classfile

import "fmt"

type ClassFile struct {

	magic	uint32
	minorVersion	uint16
	majorVersion	uint16
	constantPool	ConstantPool
	accessFlags		uint16
	thisClass		uint16
	superClass		uint16
	interfaces		[]uint16
	fields			[]*MemberInfo
	methods			[]*MemberInfo
	attributes 		[]AttributeInfo
}

/*
decode []byte data to ClassFile struct
*/
func Parese(classData []byte) (cf *ClassFile, err error) {

	//作用：相当于catch异常
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v",r)
			}
		}
	}()

	/*
	three different define types:
	var cr ClassReader
	var cr *ClassReader = &ClassReader{classData}
	var cr *ClassReader = new(User)
	*/
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

/*
读取class文件是按顺序严格执行的
 */
func (self *ClassFile) read(reader *ClassReader) {

	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	//accessFlags 是一个16位的，bitmask，ACC_PUBLIC,ACC_PRIVATE...
	self.accessFlags = reader.readUint16()
	//当前类在常量池中的索引
	self.thisClass = reader.readUint16()
	//当前类父类在常量池中的索引，只有java.lang.Object的父类这里索引位0,因为没有父类
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	//读取字段表
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

/*
01 check magic number (魔数，Java-class 文件的魔数： 0xCAFEBABE)
 */
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {

	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		//if magic do not equal to 0xCAFEBABE, panic
		panic("java.lang.ClassFormatError: magic!")
	}
}

/*
02 check version (Java 版本号 JAVA 8版本号：52.0)
 	特定的Java虚拟机实现只能支持版本号在某个范围内的class文件，oracle的实现是向后兼容，java 8 支持的版本号为 45.0 ~ 52.0
 */
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {

	//次版本号 J2SE 1.2之后，次版本号没用，都为0
	self.minorVersion = reader.readUint16()
	//主版本号
	self.majorVersion = reader.readUint16()

	switch self.majorVersion {
	//这里拿出45是因为，J2SE的主版本号是45，所以他由次版本号，不过不用考虑
	case 45:
		return
	case 46,47,48,49,50,51,52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

/*
getter
 */
func (self *ClassFile) MinorVersion() uint16 {

	return self.minorVersion
}

/*
getter
 */
func (self *ClassFile) MajorVersion() uint16 {

	return self.majorVersion
}

/*
getter
 */
func (self *ClassFile) ConstantPool() ConstantPool {

}

/*
getter
 */
func (self *ClassFile) AccessFlags() uint16 {

	return self.accessFlags
}

/*
getter
 */
func (self *ClassFile) Fields() []*MemberInfo {

	return self.fields
}

/*
getter
 */
func (self *ClassFile) Methods() []*MemberInfo {

	return self.methods
}

/*
getter
 */
func (self *ClassFile) ClassName() string {

	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {

	if self.superClass > 0 {

		return self.constantPool.getClassName(self.superClass)
	}
	return "" //只有java.lang.Object没有父类
}

func (self *ClassFile) InterfaceNames() []string {

	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces{

		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}