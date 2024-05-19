package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

const (
	sourceDir = "./pdfdir"
	noProcessRoutines = 5
)

func main() {
	var wg sync.WaitGroup
	var processWg sync.WaitGroup

	paths := make(chan string)
	output := make(chan string)

	fmt.Println("Starting find routine...")
	wg.Add(1)
	go findPDFs(sourceDir, paths, &wg)

	fmt.Printf("Starting %d process routines...\n", noProcessRoutines)
	processWg.Add(noProcessRoutines)
	for i := 0; i < noProcessRoutines; i++ {
		wg.Add(1)
		go processPDFs(paths, output, &wg, &processWg)
	}

	// Start a goroutine to close the output channel once all processing is done
	go func() {
		processWg.Wait()
		close(output)
		fmt.Println("All processing done.")
	}()

	fmt.Println("Starting collect output routine...")
	wg.Add(1)
	results := make([]string, 0)
	go collectOutput(output, &results, &wg)

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the results
	for i, result := range results {
		fmt.Printf("Result #%d: %s\n", i+1, result)
	}
}

func findPDFs(dir string, paths chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".pdf" {
			fmt.Println("Found PDF:", path)
			paths <- path
		}

		return nil
	})

	close(paths)
	fmt.Println("Finding PDFs done.")
}

func processPDFs(paths, output chan string, wg, processWg *sync.WaitGroup) {
	defer wg.Done()
	defer processWg.Done()  // Decrement the counter when the goroutine completes

	for path := range paths {
		// Process each PDF file
		// For this example, simply send the file path to the output channel
		fmt.Println("Processing PDF:", path)
		output <- path
	}
}

func collectOutput(output chan string, results *[]string, wg *sync.WaitGroup) {
	defer wg.Done()

	for path := range output {
		*results = append(*results, path)
	}
	fmt.Println("Collecting output done.")
}
