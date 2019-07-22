package classfile

/*
CONSTANT_Utf8_info{
	u1	tag;	//1
	u2	length;
	u1	bytes[length];
}
*/
type ConstantUtf8Info struct {

	str string
}

/*
需要注意的一点，字符串在 class文件中是以MUTF-8方式编码的，并不是标准的utf-8编码，两者大致相同，但不兼容
主要差别：
（1）null字符（代码点U+0000）会被编码成2字节：0xC0、0x80
（2）补充字符（Supplementary Characters,代码点大于U+FFFF的Unicode字符）
 */
func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {

	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

/*
简易版，完整版请看源代码，
如果字符串中不包含null字符或补充字符是可以用简易版的
 */
func decodeMUTF8(bytes []byte) string {

	return string(bytes)
}

