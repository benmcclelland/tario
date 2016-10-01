# tario
[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/benmcclelland/tario) [![build](https://img.shields.io/travis/benmcclelland/tario.svg?style=flat)](https://travis-ci.org/benmcclelland/tario)
reads and writes to/from tar io.Reader, io.Writer, io.WriteAt turn into data stream

golang library for reading tar files/streams offsets as if they were just the file at that offsets
also does some minimal validation of the tar header at the offset