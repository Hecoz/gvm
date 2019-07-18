package main

import (
	"archive/zip"
	"fmt"
	"path/filepath"
)

func main()  {

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
