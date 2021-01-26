package parser

import (
	"strconv"

	"github.com/rvalessandro/mf-backend/utils/errors"
)

func ParseID(IDParam string) (int64, *errors.RestErr) {
	ID, err := strconv.ParseInt(IDParam, 10, 64)
	if err != nil {
		return 0, errors.NewErrBadRequest("invalid id param")
	}

	return ID, nil
}
