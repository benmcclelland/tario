package tario

import (
	"archive/tar"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

// Validate gets and validates the next header within the tarfile
func Validate(r io.Reader) (*tar.Header, error) {
	tr := tar.NewReader(r)
	// Next will find the next header and read it in
	// this skips all meta-headers like long names, etc
	hdr, err := tr.Next()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("read header"))
	}
	return hdr, nil
}
