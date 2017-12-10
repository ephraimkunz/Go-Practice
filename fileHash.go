package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"strings"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `fileName`")
var memprofile = flag.String("memprofile", "", "write memory profile to `fileName`")

/*
Create SHA-1 hash of files in passed in directory
Write filenames / hashes to disk
*/

type shaResult struct {
	path   string
	digest []byte
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s: fileHash [args] filepath\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 1 {
		usage()
		log.Fatal()
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	pathList := make([]string, 0)

	root := flag.Arg(0)

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

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}
