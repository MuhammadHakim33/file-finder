package main

import (
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dlclark/regexp2"
)

const ROOT string = `C:\Users\Hakim\Documents`
const PATTERN string = `^BEFORE`

var files []string

func main() {

	start := time.Now()
	log.Println("Start")

	if err := filepath.WalkDir(ROOT, walkFunc); err != nil {
		log.Println(err)
	}

	end := time.Since(start)
	log.Println("Time Execution : ", end)

	for _, file := range files {
		log.Println(file)
	}
}

func walkFunc(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
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
	dir, err := os.Open(path)
	if err != nil {
		log.Printf("error opening directory: %v", err)
		return true
	}
	defer dir.Close()

	_, err = dir.Readdirnames(1)

	// if dir is empty, Readdirnames return this error [io.EOF]
	if err == io.EOF {
		return true
	}

	if err != nil {
		log.Printf("error reading directory: %v", err.Error())
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
