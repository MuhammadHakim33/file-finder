package internal

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Finder struct {
	rgx *regexp.Regexp
	dir string
}

func New(keyword, dir, ext string) (*Finder, error) {
	expr := "(?i)" + regexp.QuoteMeta(keyword)

	if ext != "" {
		if !strings.HasPrefix(ext, ".") {
			ext = "." + ext
		}
		expr += regexp.QuoteMeta(ext) + "$"
	}

	// init regex
	rgx, err := regexp.Compile(expr)
	if err != nil {
		return nil, err
	}

	// init finder
	finder := &Finder{
		rgx: rgx,
		dir: dir,
	}

	// if dir is empty, set value to current directory
	if dir == "" {
		ex, err := os.Executable()
		if err != nil {
			return nil, err
		}
		finder.dir = filepath.Dir(ex)
	}

	return finder, nil
}

func (f *Finder) Find() ([]*string, error) {

	var result []*string

	// access all dir & file sequential with walkdir func
	err := filepath.WalkDir(f.dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return errors.New("directory not found")
		}

		if d.IsDir() && isHiddenDir(d.Name()) {
			return filepath.SkipDir
		}

		// match regex
		if f.rgx.MatchString(d.Name()) {
			result = append(result, &path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func isHiddenDir(dirname string) bool {
	return strings.HasPrefix(dirname, ".")
}
