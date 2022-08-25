package testdata

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

var tempDir string

// copies a single file from src to dst.
// I'm using it only for testing
func cpFile(src, dst string) {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		panic(err)
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		panic(err)
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		panic(err)
	}
	if srcinfo, err = os.Stat(src); err != nil {
		panic(err)
	}

	if err = os.Chmod(dst, srcinfo.Mode()); err != nil {
		panic(err)
	}
}

// copies a whole directory recursively
// I'm using it only for testing
func cpDir(src string, dst string) {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		panic(err)
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		panic(err)
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		panic(err)
	}

	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			cpDir(srcfp, dstfp)
		} else {
			cpFile(srcfp, dstfp)
		}
	}
}

// Prepare copies './testdata' folder for test to temporary
// directory
//
// This function ensure idempotency of tests
func Prepare() {
	var err error
	if tempDir == "" {
		if tempDir, err = os.MkdirTemp("", "test-*"); err != nil {
			panic(err)
		}

		cpDir("./testdata", tempDir)
	}
}

// returns you absolute path to given path in temporary directory.
func AbsPath(path string) string {
	return filepath.Join(tempDir, path)
}

func String() string {
	return tempDir
}

// compare content o files, function returns true if they're
// matching, otherwise it returns false.
//
// The false is returned also when any error occurs.
//
// The function  isn't optimal but it serves only for
// testing purposes
func CompareFiles(a, b string) bool {
	var err error
	var aData, bData []byte

	// read 'a'
	aPath := filepath.Join(tempDir, a)
	if aData, err = ioutil.ReadFile(aPath); err != nil {
		return false
	}

	// read 'b'
	bPath := filepath.Join(tempDir, b)
	if bData, err = ioutil.ReadFile(bPath); err != nil {
		return false
	}

	// compare a with b
	if len(aData) != len(bData) {
		return false
	}
	for i, v := range aData {
		if v != bData[i] {
			return false
		}
	}

	return true
}
