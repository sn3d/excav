package dir

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Directory represents path to directory and provide basic operations
// for directories
type Directory string

func New(paths ...string) Directory {
	for i, p := range paths {
		paths[i] = strings.Trim(p, "\t\n")
	}

	path := filepath.Join(paths...)
	absPath, err := filepath.Abs(path)
	if err != nil {
		return Directory(path)
	} else {
		return Directory(absPath)
	}

	return Directory(filepath.Join(paths...))
}

// returns new created temporary directory
func Temp() Directory {
	tmpDir, err := ioutil.TempDir("", "excav-*")
	if err != nil {
		return Directory("")
	} else {
		return New(tmpDir)
	}
}

// returns absolute path of subdirectory in this
// directory. The directory may not exist, it's
// just join of names.
func (d Directory) Subdir(name string) Directory {
	n := strings.Trim(name, "\t\n ")
	return New(string(d), n)
}

// returns absolute path to file in directory. The file
// may not exist, it's just join of names
func (d Directory) File(name string) string {
	return filepath.Join(string(d), name)
}

// FindBySuffix go through directory and call 'found' function
// for every file that suffix is matching.
//
// Any low-level error is ignored and function continue without
// breaking.
func (d Directory) FindBySuffix(suffix string, found func(string)) {
	files, err := ioutil.ReadDir(string(d))
	if err != nil {
		return
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), suffix) {
			found(d.File(file.Name()))
		}
	}
}

// Remove the directory and all subdirectories and files recursively
func (d Directory) Remove() {
	os.RemoveAll(string(d))
}

// Ensure the directory is present in your system.
func (d Directory) Mkdir() {
	os.MkdirAll(string(d), 0755)
}

func (d Directory) IsNotSet() bool {
	return d == ""
}

func (d Directory) IsNotExist() bool {
	if stat, err := os.Stat(string(d)); !os.IsNotExist(err) {
		if stat.IsDir() {
			return false
		}
	}
	return true
}

func (d Directory) IsEmpty() bool {
	f, err := os.Open(string(d))
	if err != nil {
		return true
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true
	}
	return false
}

// The function removes all content but directory stay untouched.
// This function is bit different as Remove
func (d Directory) Clear() {
	files, err := ioutil.ReadDir(string(d))
	if err != nil {
		return
	}

	for _, file := range files {
		os.RemoveAll(filepath.Join(string(d), file.Name()))
	}
}

func (d Directory) String() string {
	return string(d)
}
