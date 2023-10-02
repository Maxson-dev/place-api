package serd

import (
	"strconv"
	"testing"

	"github.com/pkg/errors"
)

func Test_ParseIntID(t *testing.T) {
	// Test cases
	tests := []struct {
		strid       string
		expectedID  int64
		expectedErr error
	}{
		// Valid input case
		{"123", 123, nil},

		// Input case with value less than zero
		{"-10", 0, errors.New("id must be greater than zero")},

		// Input case with non-integer string
		{"abc", 0, errors.Wrap(&strconv.NumError{"Atoi", "abc", strconv.ErrSyntax}, "failed to parse id")},
	}

	// Run test cases
	for _, test := range tests {
		id, err := ParseIntID(test.strid)
		if id != test.expectedID {
			t.Errorf("Expected ID: %d, but got: %d", test.expectedID, id)
		}
		if err == nil && test.expectedErr != nil {
			t.Errorf("Expected error: %s, but got: nil", test.expectedErr)
		} else if err != nil && test.expectedErr == nil {
			t.Errorf("Expected no error, but got: %s", err)
		} else if err != nil && test.expectedErr != nil && err.Error() != test.expectedErr.Error() {
			t.Errorf("Expected error: %s, but got: %s", test.expectedErr, err)
		}
	}
}
