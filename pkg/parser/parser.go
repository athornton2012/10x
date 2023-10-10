package parser

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
)

func ParseCSV(filepath string) ([]map[string]string, error) {
	var records []map[string]string
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	features, err := csvReader.Read()
	if err == io.EOF {
		return nil, errors.New("Empty file provided")
	}

	if err != nil {
		return nil, err
	}

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		data := map[string]string{}

		for i, val := range record {
			data[features[i]] = val
		}

		records = append(records, data)
	}

	return records, nil
}
