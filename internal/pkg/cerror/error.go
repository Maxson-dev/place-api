// Package cerror provides custom errors
package cerror

import "github.com/pkg/errors"

var ErrNotFound = errors.New("entity is not found")
