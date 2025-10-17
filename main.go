package main

import (
	"io/fs"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

type Finder struct {
	root    string
	pattern string
	regex   *regexp.Regexp
	results []string
}

func NewFinder(root string, pattern string) (*Finder, error) {
	re, err := regexp.Compile("(?i)" + pattern)
	if err != nil {
		return nil, err
	}

	return &Finder{
		root:    root,
		pattern: pattern,
		regex:   re,
	}, nil
}

func main() {

	log.Println("Start")

	finder, err := NewFinder(`C:\Users\Hakim\Documents`, `^before`)
	if err != nil {
		log.Println(err)
		return
	}

	if err := filepath.WalkDir(finder.root, finder.walkFunc); err != nil {
		log.Println(err)
		return
	}

	for _, file := range finder.results {
		log.Println(file)
	}
}

func (f *Finder) walkFunc(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if d.IsDir() && isHiddenDir(d.Name()) {
		return filepath.SkipDir
	}

	if f.regex.MatchString(d.Name()) {
		f.results = append(f.results, path)
	}

	return nil
}

func isHiddenDir(dirname string) bool {
	return strings.HasPrefix(dirname, ".")
}
