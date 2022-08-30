package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	writeIntoFile()

	appendIntoFile()

}

func writeIntoFile() {
	// create byte array to write into file
	myData := []byte("I want to write into a file")
	// the WriteFile method returns an error if unsuccessful
	err := ioutil.WriteFile("localfile.txt", myData, 0777)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("Write successfully completed")
	readFromFile()
}

func readFromFile() {
	// read in the contents of the localfile
	data, err := ioutil.ReadFile("localfile.txt")
	// if our program was unable to read the file
	// print out the reason why it can't
	if err != nil {
		fmt.Print(err)
	}
	// if it was successful in reading the file then
	// print out the contents as a string
	fmt.Println(string(data))
}

func appendIntoFile() {
	f, err := os.OpenFile("localfile.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString("I am going to append data into the file"); err != nil {
		panic(err)

	}
	readFromFile()
}
