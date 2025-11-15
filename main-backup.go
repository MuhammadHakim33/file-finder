package main

// 1. pisahin input
// 2. pisahin main logic
// 3. pisahin output

// import (
// 	"flag"
// 	"fmt"
// 	"io/fs"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"regexp"
// 	"strings"
// 	"sync"
// )

// const MAXFILE int = 100

// type Finder struct {
// 	root        string
// 	pattern     string
// 	regex       *regexp.Regexp
// 	results     []string
// 	totalResult int
// 	mu          sync.Mutex
// 	wg          sync.WaitGroup
// }

// func main() {

// 	log.Println("Start")

// 	ex, err := os.Executable()
// 	if err != nil {
// 		panic(err)
// 	}

// 	currentDir := filepath.Dir(ex)

// 	flagDir := flag.String("dir", currentDir, "Directory To Search")
// 	flagSearch := flag.String("s", ``, "Search keyword")

// 	flag.Parse()

// 	finder, err := NewFinder(*flagDir, *flagSearch)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	if err := filepath.WalkDir(finder.root, finder.walkFunc); err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	finder.wg.Wait()

// 	for _, file := range finder.results {
// 		fmt.Println("+ ", file)
// 	}

// 	if finder.totalResult >= MAXFILE {
// 		log.Printf("Found %d matching files. Try a more specific keyword.", finder.totalResult)
// 	} else {
// 		log.Printf("Found %d matching files.", finder.totalResult)
// 	}
// }

// func NewFinder(root string, pattern string) (*Finder, error) {
// 	re, err := regexp.Compile("(?i)" + pattern)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Finder{
// 		root:    root,
// 		pattern: pattern,
// 		regex:   re,
// 	}, nil
// }

// func (f *Finder) walkFunc(path string, d fs.DirEntry, err error) error {
// 	if err != nil {
// 		return err
// 	}

// 	if d.IsDir() && isHiddenDir(d.Name()) {
// 		return filepath.SkipDir
// 	}

// 	f.wg.Add(1)
// 	go func(name, path string) {
// 		defer f.wg.Done()

// 		if f.regex.MatchString(d.Name()) {
// 			f.mu.Lock()
// 			defer f.mu.Unlock()

// 			f.totalResult++
// 			if len(f.results) <= MAXFILE {
// 				f.results = append(f.results, path)
// 			}
// 		}
// 	}(d.Name(), path)

// 	return nil
// }

// func isHiddenDir(dirname string) bool {
// 	return strings.HasPrefix(dirname, ".")
// }
