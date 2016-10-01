package tario

import "io"

// TarReader is an io.Reader on a tar file returning the
// file contained at that position of the tarfile
type TarReader struct {
	r          io.Reader
	doneHeader bool
	offset     int64
	size       int64
}

// NewReader takes an io.Reader and returns a TarReader
// the io.Reader should be position at the begining of the
// tar header for the file
func NewReader(r io.Reader) *TarReader {
	return &TarReader{r: r}
}

// Read satisfies the io.Reader interface for a TarReader
func (t *TarReader) Read(b []byte) (int, error) {
	if !t.doneHeader {
		hdr, err := Validate(t.r)
		if err != nil {
			return 0, err
		}
		t.size = hdr.Size
		t.doneHeader = true
	}

	if t.offset == t.size {
		return 0, io.EOF
	}

	if int64(len(b)) > (t.size - t.offset) {
		b = make([]byte, t.size)
	}

	i, err := t.r.Read(b)
	t.offset += int64(i)
	return i, err
}
