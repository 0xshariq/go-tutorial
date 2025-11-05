package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println("file name", fileInfo.Name())
	fmt.Println("file size", fileInfo.Size())
	fmt.Println("file permission", fileInfo.Mode())
	fmt.Println("file or folder", fileInfo.IsDir())

	// read file
	data := make([]byte, fileInfo.Size())
	count, err := f.Read(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("read %d bytes\n", count)
	fmt.Println(string(data))

	// create file
	newFile, err := os.Create("newfile.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	newFile.WriteString("this is a new file")

	// delete file
	error := os.Remove("newfile.txt")
	if error != nil {
		panic(error)
	}

	// write file
	bytes := []byte("Hello Bhai")
	_, erro := f.Write(bytes)
	if erro != nil {
		panic(erro)
	}

}
