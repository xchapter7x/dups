package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	const blockSize = 10
	dupStore := FindDuplicatesInFiles(os.Stdin, blockSize)
	fmt.Println(PrettyFormatDuplicatesString(dupStore))
}

func PrettyFormatDuplicatesString(duplicates map[string][]string) string {
	prettyString := ""
	cnt := 0
	for i, v := range duplicates {
		cnt += 1
		prettyString += fmt.Sprintln(i)
		for _, f := range v {
			prettyString += fmt.Sprintln("    ->", f)
		}
		prettyString += fmt.Sprintln("-------------------------------------------------------------")
	}
	prettyString += fmt.Sprintln(cnt, ": instances of duplication found")
	return prettyString
}

func FindDuplicatesInFiles(reader io.Reader, blockSize int) map[string][]string {
	dupStore := make(map[string][]string)

	for _, filename := range ReadFileListFromReader(reader) {
		if filename == "" {
			continue
		}

		if stat, err := os.Stat(filename); err == nil && stat.IsDir() {
			continue
		}

		lines, err := CreateStringArrayFromFile(filename)
		if err != nil {
			continue
		}

		for i, _ := range lines {
			if i+blockSize <= len(lines)-1 {
				key := strings.Replace(strings.Join(lines[i:i+blockSize], ""), " ", "", -1)
				dupStore[key] = append(dupStore[key], fmt.Sprintf("%v:%v,%v", strings.TrimSpace(filename), i, i+blockSize))
			}
		}
	}

	for k, v := range dupStore {
		if len(v) <= 1 {
			delete(dupStore, k)
		}
	}
	return dupStore
}

func CreateStringArrayFromFile(filename string) ([]string, error) {
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadFile(strings.TrimSpace(absPath))
	if err != nil {
		return nil, err
	}

	if len(string(content)) == 0 {
		return make([]string, 0), nil
	}

	lines := strings.Split(string(content), "\n")
	return lines, nil
}

func ReadFileListFromReader(fileListReader io.Reader) []string {
	reader := bufio.NewReader(fileListReader)
	filelist := make([]string, 0)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			if len(string(text)) > 0 {
				filelist = append(filelist, string(text))
			}
			break
		}
		filelist = append(filelist, string(text))
	}
	return filelist
}
