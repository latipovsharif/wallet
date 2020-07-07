package api

import (
	"bytes"
	"encoding/csv"

	"github.com/pkg/errors"
)

func toCsv(transactions []transaction) (*bytes.Buffer, error) {
	buff := &bytes.Buffer{}
	w := csv.NewWriter(buff)

	for _, trn := range transactions {
		if err := w.Write(trn.getRow()); err != nil {
			return nil, errors.Wrap(err, "cannot write transaction")
		}
	}

	return buff, nil
}
