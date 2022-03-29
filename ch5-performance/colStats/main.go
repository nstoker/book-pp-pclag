package main

import (
	"flag"
	"fmt"
	"io"
	"os"

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

	// Loop through all files adding their data to consolidate
	for _, fname := range filenames {
		f, err := os.Open(fname)
		if err != nil {
			return errors.Wrapf(err, "cannot open file %s", fname)
		}
		defer f.Close()

		data, err := csv2float(f, column)
		if err != nil {
			return err
		}

		if err := f.Close(); err != nil {
			return err
		}

		// Append the data to consolidate
		consolidate = append(consolidate, data...)
	}

	_, err := fmt.Fprintln(out, opFunc(consolidate))
	return err
}
