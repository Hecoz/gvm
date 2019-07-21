package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {

	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{absPath}
}

//extract class file from jar file or zip file
func (self *ZipEntry) readClass(className string) ([]byte,Entry,error){

	fliesReaderCloser, err := zip.OpenReader(self.absPath)
	/*
	01 check if the jar or zip file can open
	 */
	if err != nil {
		return nil,nil,err
	}

	//defer语句调用一个函数，这个函数执行会推迟，直到外围的函数返回，或者外围函数运行到最后，或者相应的goroutine panic,defer的执行是先进后出
	defer fliesReaderCloser.Close()

	/*
	02 traversing the jar or zip file to find the correct class
	 */
	for _, classFile := range fliesReaderCloser.File {

		if classFile.Name == className{

			//open classfile -- reader
			fileReaderCloser, err := classFile.Open()
			if err != nil {
				return nil, nil, err
			}

			defer fileReaderCloser.Close()
			//read data
			data, err := ioutil.ReadAll(fileReaderCloser)
			if err != nil {

				return nil,nil,err
			}
			//find class file and return
			return data, self, nil
		}
	}
	/*
	03 do not find class in jar or zip
	 */
	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {

	return self.absPath
}
