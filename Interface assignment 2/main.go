package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// file, err := os.Open("myfile.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// content, err := ioutil.ReadAll(file)

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Println(string(content))

	// Method 2
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
