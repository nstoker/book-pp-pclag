package main

import (
	"encoding/csv"
	"io"
	"strconv"

	"github.com/pkg/errors"
)

func sum(data []float64) float64 {
	sum := 0.0

	for _, v := range data {
		sum += v
	}

	return sum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

// statsFunc defines a generic statistical function
type statsFunc func(data []float64) float64

func csv2float(r io.Reader, column int) ([]float64, error) {
	// Create the csv reader used to read in data from csv files
	cr := csv.NewReader(r)

	// Adjust for a 0-based index
	column--

	// Read in all the csv data
	allData, err := cr.ReadAll()
	if err != nil {
		return nil, errors.Wrap(err, "cannot read data from file")
	}

	var data []float64

	for i, row := range allData {
		if i == 0 {
			continue
		}

		// checking number of columns in CSV file
		if len(row) <= column {
			// file does not have that many columns
			return nil, errors.Wrapf(ErrInvalidColumn, "file has only %d columns", len(row))
		}

		// try to convert to a float
		v, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			return nil, errors.Wrapf(ErrNotNumber, "%v", err)
		}

		data = append(data, v)
	}

	return data, nil
}
