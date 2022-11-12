package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"log"
	"sync"
	"time"
)

// In programming, concurrent tasks can communicate with each other and share resources. Go provides a way for bidirectional communication between two goroutines through channels.
type ZipWriter struct {
	buffer bytes.Buffer
	zip    zip.Writer
}

func writeChannel(wg *sync.WaitGroup, limitChannel chan int, stop int) {
	defer wg.Done()
	for i := 1; i <= stop; i++ {
		limitChannel <- i
	}
}

func readChannel(wg *sync.WaitGroup, limitChannel chan int, stop int) {
	defer wg.Done()
	for i := 1; i <= stop; i++ {
		fmt.Println(<-limitChannel)
		fmt.Println(time.Now().Format(time.RFC1123Z))
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	limitChannel := make(chan int)
	defer close(limitChannel)
	go writeChannel(&wg, limitChannel, 3)
	go readChannel(&wg, limitChannel, 3)

	NewZipWriter()
	wg.Wait()
}

func NewZipWriter() {
	writer := &ZipWriter{}
	writer.buffer = bytes.Buffer{}
	writer.zip = *zip.NewWriter(&writer.buffer)
	// fmt.Println(writer.zip)
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		f, err := writer.zip.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("File created successfully")
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	err := writer.zip.Close()
	if err != nil {
		log.Fatal(err)
	}

}
