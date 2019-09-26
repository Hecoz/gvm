package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {

	absDir string
}

//以new开头作为构造函数，创建结构体实例
func newDirEntry(path string) *DirEntry{

	absDir,err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	//相当与 new 了一个DirEntry，并返回其地址,因为函数声明中，返回值是指针
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte,Entry,error){

	//join to a full path
	fileName := filepath.Join(self.absDir,className)
	//read file
	data,err := ioutil.ReadFile(fileName)
	return data,self,err
}

func (self *DirEntry) String() string {

	return self.absDir
}