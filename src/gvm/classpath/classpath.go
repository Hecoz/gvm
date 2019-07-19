package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {

	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {

	cp := &Classpath{}
	//01 首先解析 启动类路径 和 扩展类路径
	cp.parseBootAndExtClasspath(jreOption)
	//02 其次解析 用户类路径
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {

	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string{

	return self.userClasspath.String()
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string)  {

	//01 解析 JRE 路径
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir,"lib","*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext","*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func (self *Classpath) parseUserClasspath(cpOption string) {

	if cpOption == ""{
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

/**
获取JRE路径
01 首先检查用户是否通过-Xjre自定义JRE路径，如果重新设置则直接返回
02 如果用户没有设置，检查当前的路径中是否存在JRE文件，如果存在直接返回当前路径下的JRE
03 如果没有，检查系统环境变量中JAVA_HOME下的JRE路径
 */
func getJreDir(jreOption string) string {

	//01 check if -Xjre option is not null and jreOption exists
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	//02 check if ./jre exists in current dir
	if exists("./jre"){
		return "./jre"
	}
	//03 check if jre exists in JAVA_HOME
	if jh := os.Getenv("JAVA_HOME"); jh != ""{
		return filepath.Join(jh,"jre")
	}
	//throw error can not find jre folder
	panic("Can not find jre folder!")
}

//check if dir exists
func exists(path string) bool {

	//read file from path , check if error equals nil, if not ,may not exists
	if _, err := os.Stat(path); err != nil{

		if os.IsNotExist(err){
			return false
		}
	}
	return true
}