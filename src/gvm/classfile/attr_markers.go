package classfile


/*
Deprecated_attribute{
	u2	attribute_name_index;
	u4	attribute_length;	//0
}
 */
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
Synthetic_attribute{
	u2	attribute_name_index;
	u4	attribute_length;
}
 */
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {}

func (self *MarkerAttribute) readInfo(reader *ClassReader)  {

	//read nothing
	//these two attributes has no data, so readInfo method is empty
}