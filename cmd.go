package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct{

	helpFlag 	bool 		// print help content
	versionFlag bool		// print gvm version
	cpOption	string		// print class path option
	XjreOption	string		// print jre option
	class		string		// print class
	args		[]string	// print args
}

//解析命令行参数
func parseCmd() *Cmd {

	cmd := &Cmd{}

	//改变 flag 默认的 Usage
	flag.Usage = printUsage
	//绑定到变量 cmd.helpFlag， 使用参数为 -help 或 -? ， 默认值为 false ， 如果没有重置 flag 默认的Usage，则输出这里的usage
	flag.BoolVar(&cmd.helpFlag, "help",false,"print help message")
	flag.BoolVar(&cmd.helpFlag,"?",false,"print help message")
	flag.BoolVar(&cmd.versionFlag,"version",false,"print version and exit")
	flag.StringVar(&cmd.cpOption,"classpath","","classpath")
	flag.StringVar(&cmd.cpOption,"cp","","classpath")
	flag.StringVar(&cmd.XjreOption,"Xjre","","path to jre")
	flag.Parse()

	//这里，如果没有用参数指定，非参数选项，则输入的就为 args,
	//例如 gvm java.lang.Object => cmd.class = java.lang.Object
	//    gvm java.lang.Object 1 2 3 => cmd.class = java.lang.Object , cmd.args = 1,2,3
	args := flag.Args()

	if len(args) > 0 {

		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

/*
输出命令输入错误时候的提示信息
 */
func printUsage(){
	fmt.Printf("Usage: %s [-options] class [args...]\n",os.Args[0])
}

