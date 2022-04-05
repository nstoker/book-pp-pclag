package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"

	"github.com/pkg/errors"
)

func main() {
	// Verify and parse arguments
	op := flag.String("op", "sum", "Operation to be executed")
	column := flag.Int("col", 1, "CSV column on which to commence operation")
	flag.Parse()

	if err := run(flag.Args(), *op, *column, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filenames []string, op string, column int, out io.Writer) error {
	var opFunc statsFunc

	if len(filenames) == 0 {
		return ErrNoFiles
	}

	if column < 1 {
		return errors.Wrapf(ErrInvalidColumn, "%d", column)
	}

	// Validate the operation and define the opFunc accordingly
	switch op {
	case "sum":
		opFunc = sum
	case "avg":
		opFunc = avg
	default:
		return errors.Wrapf(ErrInvalidOperation, "%s", op)
	}

	consolidate := make([]float64, 0)

	// Create the channel to receive results or errors of operations
	filesCh := make(chan string)
	resCh := make(chan []float64)
	errCh := make(chan error)
	doneCh := make(chan struct{})

	wg := sync.WaitGroup{}

	// Loop through all files and send them through the channel so
	// each one will be processed when a worker is available
	go func() {
		defer close(filesCh)

		for _, fname := range filenames {
			filesCh <- fname
		}
	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {

			defer wg.Done()
			for fname := range filesCh {
				// open the file for reading
				f, err := os.Open(fname)
				if err != nil {
					errCh <- errors.Wrapf(err, "cannot open file %s", fname)
					return
				}

				// Parse the CSV into a slice of float64 numbers
				data, err := csv2float(f, column)
				if err != nil {
					errCh <- err
				}

				if err := f.Close(); err != nil {
					errCh <- err
				}

				resCh <- data
			}
		}()
	}
	go func() {
		wg.Wait()
		close(doneCh)
	}()

	for {
		select {
		case err := <-errCh:
			return err
		case data := <-resCh:
			consolidate = append(consolidate, data...)
		case <-doneCh:
			_, err := fmt.Fprintln(out, opFunc(consolidate))
			return err
		}
	}
}
