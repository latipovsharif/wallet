package api

import (
	"bytes"
	"encoding/csv"

	"github.com/pkg/errors"
)

func toCsv(transactions []transaction) (*bytes.Buffer, error) {
	buff := &bytes.Buffer{}
	w := csv.NewWriter(buff)

	if err := w.Write(getCSVHeader()); err != nil {
		return nil, errors.Wrap(err, "cannot write csv file headers")
	}

	for _, trn := range transactions {
		if err := w.Write(trn.getCSVRow()); err != nil {
			return nil, errors.Wrap(err, "cannot write transaction")
		}
	}

	w.Flush()

	return buff, nil
}
