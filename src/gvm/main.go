package main

import (
	"fmt"
	"gvm/classfile"
	"gvm/classpath"
	"strings"
)

func main() {

	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == ""{
		printUsage()
	} else {
		startJVM(cmd)
	}
}

//test cmd: gvm -Xjre $JAVA_HOME/jre java.lang.Object
//test cmd: gvm "./javaClasses/CUT"
func startJVM(cmd *Cmd){

	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)

	fmt.Printf("classpath:%v class:%v args:%v\n",cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".","/",-1)

	fmt.Println("cp:",cp)
	fmt.Println("className:", className)

	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)


	//classData, _, err := cp.ReadClass(className)
	//
	//if err != nil {
	//
	//	fmt.Printf("Could not find or load main class %s\n",cmd.class)
	//	return
	//}
	//
	//fmt.Printf("class data: %v\n",classData)
	//fmt.Printf("classpath:%s class:%s args:%v\n",cmd.cpOption,cmd.class,cmd.args)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {

	classData, _, err := cp.ReadClass(className)
	if err != nil{
		panic(err)
	}

	cf, err := classfile.Parese(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile)  {

	fmt.Printf("version: %v. %v\n",cf.MajorVersion(),cf.MinorVersion())
	fmt.Printf("constants count: %v\n",len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields(){
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("mehtods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods(){

		fmt.Printf("  %s\n", m.Name())
	}
}
