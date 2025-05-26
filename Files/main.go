package main

import (
	"fmt"
	"os"
)

func main() {
	//f, err := os.Open("./example.txt")
	//if err != nil {
	//	// log the error
	//	panic(err)
	//}
	//
	//fileInfo, err := f.Stat()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("FileInfo:", fileInfo.Name())       //Name of the Files
	//fmt.Println("Is directory:", fileInfo.IsDir())  //It's a directory or not
	//fmt.Println("file size:", fileInfo.Size())      //File size
	//fmt.Println("Is directory:", fileInfo.Mode())   //File permission
	//fmt.Println("Modified at:", fileInfo.ModTime()) //Modified at

	// Read file (we can read file like below)
	//f, err := os.Open("./example.txt")
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer f.Close()
	//
	//buff := make([]byte, 20)
	//d, err := f.Read(buff)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for i := 0; i < len(buff); i++ {
	//	println("Data", d, string(buff[i]))
	//}

	// Read file
	//file, err := os.ReadFile("example.txt")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(file))
	//
	//// Read Folder
	//// if we pass ("../") it will go one step back
	//dir, err := os.Open("../")
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer dir.Close()
	//
	//fileInfo, err := dir.Readdir(-1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, file := range fileInfo {
	//	fmt.Println(file.Name())
	//}

	// Create A file
	//file, err := os.Create("email.txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()

	//file.WriteString("Hi Everyone")
	//file.WriteString("\nLove you all")

	//bytes := []byte("Hello Golang")
	//file.Write(bytes)

	// Read and Write to another file (Streaming fashion)
	//sourceFile, err := os.Open("example.txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer sourceFile.Close()
	//
	//destFile, err := os.Create("example2.txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer destFile.Close()
	//reader := bufio.NewReader(sourceFile)
	//writer := bufio.NewWriter(destFile)
	//
	//for {
	//	b, err := reader.ReadByte()
	//	if err != nil {
	//		if err.Error() != "EOF" {
	//			panic(err)
	//		}
	//
	//		break
	//	}
	//	e := writer.WriteByte(b)
	//	if e != nil {
	//		panic(e)
	//	}
	//}
	//
	//writer.Flush()
	//
	//fmt.Println("Return to new file")

	// Delete A file
	err := os.Remove("email.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Files Deleted Successfully")
}
