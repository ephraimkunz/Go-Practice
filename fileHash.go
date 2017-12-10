package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

/*
Create SHA-1 hash of files in passed in directory
Write filenames / hashes to disk
*/

type shaResult struct {
	path   string
	digest []byte
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./test <root filepath>")
		os.Exit(1)
	}

	pathList := make([]string, 0)

	root := os.Args[1]
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && !strings.Contains(path, ".git") {
			pathList = append(pathList, path)
		}
		return nil
	})

	fmt.Printf("%d paths to process\n", len(pathList))
	filepathChan := make(chan string, len(pathList))
	resultChan := make(chan shaResult, len(pathList))

	for _, path := range pathList {
		filepathChan <- path // Load filepaths onto a channel for processing by go routines
	}
	close(filepathChan)

	const NumGoRoutines = 7
	fmt.Printf("Spinning up %d goroutines to process these\n", NumGoRoutines)

	// Keep track of when all finish goroutines finishs
	done := make(chan bool, NumGoRoutines)

	for i := 0; i < NumGoRoutines; i++ {
		go func(id int) {
			var numProcessed uint64
			for path := range filepathChan {
				dat, err := ioutil.ReadFile(path)
				if err == nil {
					h := sha256.New()
					h.Write(dat)
					bs := h.Sum(nil)
					result := shaResult{path, bs}
					numProcessed++
					resultChan <- result
				}
			}
			fmt.Printf("Chan %d processed %d items\n", id, numProcessed)
			done <- true
		}(i)
	}

	f, _ := os.Create("./digest_dump")
	defer f.Close()

	for i := 0; i < len(pathList); i++ {
		item := <-resultChan
		fmt.Fprintf(f, "%s\n\t%x\n\n", item.path, item.digest)
	}

	for i := 0; i < NumGoRoutines; i++ {
		<-done
	}
}
