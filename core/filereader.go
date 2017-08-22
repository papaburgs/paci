package core

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var ErrMissingDatafile = errors.New("Missing file")

func ReadAll(path string) ([]Ingredient, []Item) {
	cc.log.Debugf("parsing path: %s", path)
	err := filepath.Walk(path, walker)
	if err != nil {
		panic(err)
	}
	return []Ingredient{}, []Item{}
}

func walker(path string, info os.FileInfo, err error) error {
	var (
		blob []byte
	)
	cc.log.Debug("another iteration -----------------------------------------")
	if err != nil {
		cc.log.Warnf("error from walk on %s:\n%s", path, err)
		return err
	}
	if info.IsDir() {
		return nil
	}
	blob, err = getblob(path)
	return nil
}

func getblob(path string) ([]byte, error) {
	var (
		err    error
		infile *os.File
		blob   []byte
	)
	infile, err = os.Open(path)
	if err != nil {
		cc.log.Error("01084 ", err)
		return []byte{}, ErrMissingDatafile
	}

	// close file when done
	defer func(f *os.File) {
		e := f.Close()
		if e != nil {
			cc.log.Error(e)
		}
	}(infile)

	blob, err = ioutil.ReadAll(infile)
	if err != nil {
		cc.log.Error("75291 ", err)
		return []byte{}, err
	}

	return blob, err
}
