package serd

import (
	"strconv"

	"github.com/pkg/errors"
)

func ParseIntID(strid string) (int64, error) {
	i, err := strconv.Atoi(strid)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse id")
	}
	if i <= 0 {
		return 0, errors.New("id must be greater than zero")
	}
	return int64(i), nil
}
