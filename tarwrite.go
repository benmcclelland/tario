package tario

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// TARHEADERSIZE is size of a tar header in bytes
const TARHEADERSIZE = 512

// TarWriter is an io.WriterAt being given a tar file
// and writing out the resulting file contained at that
// position of the tarfile
type TarWriter struct {
	wa       io.WriterAt
	w        io.WriteCloser
	notBegin bool
}

// NewWriter takes an io.WriterAt and returns a TarWriter
func NewWriter(wa io.WriterAt) *TarWriter {
	return &TarWriter{wa: wa}
}

func NewFileWriter() *TarWriter {
	return &TarWriter{}
}

// WriteAt satisfies the io.WriterAt interface for a TarWriter
// we assume that the write at offset 0 will include the entire tar header
// and that no other offset write overlaps the header
func (t *TarWriter) WriteAt(p []byte, off int64) (n int, err error) {
	if off == 0 {
		_, err := Validate(bytes.NewBuffer(p))
		if err != nil {
			return 0, err
		}
		n, err := t.wa.WriteAt(p[512:], 0)
		if err != nil {
			return 0, err
		}
		return n + 512, nil
	}
	return t.wa.WriteAt(p, (off - 512))
}

func (t *TarWriter) Write(p []byte) (n int, err error) {
	if t.notBegin == false {
		t.notBegin = true
		hdr, err := Validate(bytes.NewBuffer(p))
		if err != nil {
			return 0, err
		}
		//t.w, err = os.Create(hdr.Name)
		fmt.Println("TARHEADER:", hdr.Name, hdr.Size)
		t.w, err = os.Create(hdr.Name)
		if err != nil {
			log.Fatal(err)
		}

		n, err := t.w.Write(p[512:])
		if err != nil {
			return 0, err
		}
		return n + 512, nil
	}
	return t.w.Write(p)
}

func (t *TarWriter) Close() error {
	return t.w.Close()
}
