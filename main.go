package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("dummyfile - copyright (c) 2021 by Laurenz Wagner")
	fmt.Println("------------------------------------------------------------------")

	// create 50 1kb files
	// create 10 1MB files
	// create 1 10 MB, 1 20 MB

	sizeMB := 10 // first file. next file size is double (10 MB, 20 MB, 40 MB, 80 MB, 160 MB )
	sizekB := sizeMB * 1000 * 1000

	n := 1
	for n <= 8 {
		//fmt.Println(n)
		createTestFile("testfile"+strconv.Itoa(n), dataBySize(sizekB))
		sizekB = sizekB * 2
		n = n + 1
	}

	fmt.Println("---------------------------- DONE --------------------------------")

}

func createTestFile(fileName string, data []byte) {

	f, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	defer f.Close()
	size, err := f.Write(data) // Write it to the file
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	sizeWrittenInMB := size / 1000 / 1000
	fmt.Println(strconv.Itoa(sizeWrittenInMB) + " MB written to " + fileName)

}

func dataBySize(sizeInMegaByte int) (data []byte) {
	//data := make([]byte, int(1e7), int(1e7)) // Initialize an empty byte slice
	data = make([]byte, sizeInMegaByte, sizeInMegaByte) // Initialize an empty byte slice

	return data
}
