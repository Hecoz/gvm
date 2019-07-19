package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"path/filepath"
)

func main1()  {

	//testReader()
	testFlag()
}

func testFlag()  {

	//test cmd -help -version -classpath "wo" 1 2 3
	//bool 变量后面不能接参数
	var helpFlag bool
	var versionFlag bool
	var cpOption string

	//如果没有重置 flag 默认的Usage，则输出这里的usage
	//flag.Usage = prints
	flag.BoolVar(&helpFlag, "help",false,"print help message")
	flag.BoolVar(&helpFlag,"?",false,"print help message")
	flag.BoolVar(&versionFlag,"version",false,"print version and exit")
	flag.StringVar(&cpOption,"classpath","","classpath")
	flag.StringVar(&cpOption,"cp","","classpath")
	flag.Parse()

	args := flag.Args()

	fmt.Println("helpFlag: ",helpFlag)
	fmt.Println("versionFlag: ",versionFlag)
	fmt.Println("cpOption: ",cpOption)
	fmt.Println("args: ",args)
}

//func prints(){
//	fmt.Printf("Usage: %s [-options] class [args...]\n",os.Args[0])
//}



//测试读取JAR包
func testReader()  {

	path := "./src/gvm/lib"

	abspath,_ := filepath.Abs(path)

	fmt.Println(abspath)

	filename := filepath.Join(abspath,"asm.jar")

	fmt.Println(filename)

	readerCloser, err := zip.OpenReader(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer readerCloser.Close()

	for _, f := range readerCloser.File {

		fmt.Println(f.Name)
	}
}
