package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	fmt.Println("dummyfile - copyright (c) 2021 by Laurenz Wagner")
	fmt.Println("------------------------------------------------------------------")

	// Subcommands
	tinyFilesCommand := flag.NewFlagSet("tiny", flag.ExitOnError)
	bigFilesCommand := flag.NewFlagSet("big", flag.ExitOnError)

	// Tiny subcommand flag pointer
	tinyFileSizeMin := tinyFilesCommand.String("sizeMin", "40", "dummyfile tiny <40> (Minimum file size in kB)")
	tinyFileSizeMax := tinyFilesCommand.String("sizeMax", "40", "Maximum file size in kB")
	tinyFileCount := tinyFilesCommand.String("fileCount", "5000", "Count of files to generate")

	// Big subcommand flag pointer
	bigFileSizeStart := bigFilesCommand.String("sizeStart", "10", "1st file size in MB")
	bigFileCount := bigFilesCommand.String("fileCount", "8", "Count of files to generate - file size is doubled every time")

	// Verify that a subcommand has been provided
	// os.Arg[0] is the main command
	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("ERROR: tiny or big subcommand is required!")
		fmt.Println("Eg.: dummyfile tiny .... or dummyfile big ....")
		os.Exit(1)
	}

	// Switch on the subcommand
	// Parse the flags for appropriate FlagSet
	// FlagSet.Parse() requires a set of arguments to parse as input
	// os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
	switch os.Args[1] {
	case "tiny":
		tinyFilesCommand.Parse(os.Args[2:])
	case "big":
		bigFilesCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check which subcommand was Parsed using the FlagSet.Parsed() function. Handle each case accordingly.
	// FlagSet.Parse() will evaluate to false if no flags were parsed (i.e. the user did not provide any flags)
	if tinyFilesCommand.Parsed() {
		// Required Flags
		if *tinyFileSizeMin == "" {
			tinyFilesCommand.PrintDefaults()
			os.Exit(1)
		}
		if *tinyFileSizeMax == "" {
			tinyFilesCommand.PrintDefaults()
			os.Exit(1)
		}
		if *tinyFileCount == "" {
			tinyFilesCommand.PrintDefaults()
			os.Exit(1)
		}

		// we got some values

		sizeMin, err := strconv.Atoi(*tinyFileSizeMin)
		if err != nil {
			tinyFilesCommand.PrintDefaults()
			os.Exit(1)
		}
		sizeMax, err := strconv.Atoi(*tinyFileSizeMax)
		if err != nil {
			tinyFilesCommand.PrintDefaults()
			os.Exit(1)
		}
		if sizeMin > sizeMax {
			tinyFilesCommand.PrintDefaults()
			os.Exit(1)
		}
		fileCount, err := strconv.Atoi(*tinyFileCount)
		if err != nil {
			tinyFilesCommand.PrintDefaults()
			os.Exit(1)
		}

		n := 1
		sizekB := sizeMin * 1000
		for n <= fileCount {

			if sizeMin != sizeMax {
				sizekB = (rand.Intn(sizeMax-sizeMin) + sizeMin) * 1000
			}

			createTestFile("tinyfile"+strconv.Itoa(n), dataBySize(sizekB))

			n = n + 1
		}

	}

	if bigFilesCommand.Parsed() {
		// Required Flags
		if *bigFileSizeStart == "" {
			bigFilesCommand.PrintDefaults()
			os.Exit(1)
		}
		if *bigFileCount == "" {
			bigFilesCommand.PrintDefaults()
			os.Exit(1)
		}

		//Print
		fmt.Printf("bigFileSizeStart: %s, bigFileCount: %s\n",
			*bigFileSizeStart,
			*bigFileCount,
		)

		// we got some values
		//sizeMB := 10 // first file. next file size is doubled in size ... (10 MB, 20 MB, 40 MB, 80 MB, 160 MB )
		sizeMB, err := strconv.Atoi(*bigFileSizeStart)
		if err != nil {
			bigFilesCommand.PrintDefaults()
			os.Exit(1)
		}

		sizekB := sizeMB * 1000 * 1000

		countOfFiles, err := strconv.Atoi(*bigFileCount)
		if err != nil {
			bigFilesCommand.PrintDefaults()
			os.Exit(1)
		}

		n := 1
		for n <= countOfFiles {
			//fmt.Println(n)
			createTestFile("bigfile"+strconv.Itoa(n), dataBySize(sizekB))
			sizekB = sizekB * 2
			n = n + 1
		}

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
