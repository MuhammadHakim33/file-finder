package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dlclark/regexp2"
)

const ROOT string = `C:\Users\Hakim\Documents\test`
const PATTERN string = `^BEFORE`

var files []string

func main() {

	if err := filepath.WalkDir(ROOT, walkFunc); err != nil {
		log.Println(err)
	}

	for _, file := range files {
		log.Println(file)
	}
}

func walkFunc(path string, d fs.DirEntry, err error) error {
	if err != nil {
		log.Printf("error visit directory : %v", err)
		return nil
	}

	if d.IsDir() && isHiddenDir(d.Name()) {
		return filepath.SkipDir
	}

	if d.IsDir() && isEmptyDir(path) {
		return filepath.SkipDir
	}

	if isMatchPattern(PATTERN, d.Name()) {
		files = append(files, path)
	}

	return nil
}

func isHiddenDir(dirname string) bool {
	return strings.HasPrefix(dirname, ".")
}

func isEmptyDir(path string) bool {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Printf("error visit directory : %v", err)
		return true
	}

	if len(entries) == 0 {
		return true
	}

	return false
}

func isMatchPattern(pattern string, name string) bool {
	reg, err := regexp2.Compile(pattern, regexp2.IgnoreCase)
	if err != nil {
		log.Printf("error match pattern regex : %v", err)
		return false
	}

	if isMatch, _ := reg.MatchString(name); isMatch {
		return true
	}

	return false
}
