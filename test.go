package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"path/filepath"
)

func main1()  {

	//testReader()
	//testFlag()
	//testBiBao()
	testPanic()
}


//TEST SSH
//defer 的执行是先进后出
func testPanic()  {

	defer func(){
		if err := recover() ; err != nil {
			fmt.Println(err)
		}
	}()
	defer func(){
		panic("three")
	}()
	defer func(){
		panic("two")
	}()
	panic("one")
}

//测试闭包
func testBiBao()  {

	add_func := add(1,2)
	fmt.Println(add_func())
	fmt.Println(add_func())
	fmt.Println(add_func())
}
//闭包函数
/*
func() {
    //func body
}()     //花括号后加()表示函数调用，此处声明时为指定参数列表，
        //故调用执行时也不需要传参
 */
// 函数名： add  参数： x1,x2 int  返回值: func()(int,int)
func add(x1, x2 int) func()(int,int)  {
	i := 0
	return func() (int,int){
		i++
		return i,x1+x2
	}
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
