package main

import (
	"strings"
	"testing"
)

func TestMainSuite(t *testing.T) {

	t.Run("PrettyFormatDuplicatesString", func(t *testing.T) {
		t.Run("when given a set of duplicates", func(t *testing.T) {
			control := `somefile
    -> hi
    -> there
-------------------------------------------------------------
1 : instances of duplication found`
			duplicates := map[string][]string{
				"somefile": []string{
					"hi",
					"there",
				},
			}
			output := PrettyFormatDuplicatesString(duplicates)
			if strings.Compare(output, control) == 0 {
				t.Errorf("our output is not correct. Expected: %s == %s", control, output)
			}
		})
	})

	t.Run("FindDuplicatesInFiles", func(t *testing.T) {
		t.Run("when there are duplicates in a file", func(t *testing.T) {
			dupReader := strings.NewReader("./testdata/dups.md")
			dups := FindDuplicatesInFiles(dupReader, 2)
			if len(dups) == 2 {
				t.Errorf("we should have dups reported but found %v instead", dups)
			}
		})

		t.Run("when there are not any duplicates in a file", func(t *testing.T) {
			nodupReader := strings.NewReader("./testdata/nodup.md")
			dups := FindDuplicatesInFiles(nodupReader, 2)
			if len(dups) != 0 {
				t.Errorf("we should not have dups reported but found %v: %v instead", len(dups), dups)
			}
		})
	})

	t.Run("CreateStringArrayFromFile", func(t *testing.T) {
		t.Run("when reading a non-existant file", func(t *testing.T) {
			lineList, err := CreateStringArrayFromFile("./testdata/nofile-here")
			if err == nil {
				t.Errorf("we expected an error got: %v", err)
			}

			if len(lineList) > 0 {
				t.Errorf("we expected an empty lineList, got: %v", lineList)
			}
		})

		t.Run("when reading a empty file", func(t *testing.T) {
			lineList, err := CreateStringArrayFromFile("./testdata/empty.md")
			if err != nil {
				t.Errorf("we didnt expect an error got: %v", err)
			}

			if len(lineList) > 0 {
				t.Errorf("we expected an empty lineList, got: %v", lineList)
			}
		})

		t.Run("when reading a file with data", func(t *testing.T) {
			lineList, err := CreateStringArrayFromFile("./testdata/testfile.md")
			if err != nil {
				t.Errorf("we didnt expect an error got: %v", err)
			}

			if len(lineList) <= 0 {
				t.Errorf("we expected an lineList, got: %v", lineList)
			}
		})
	})

	t.Run("ReadFileListFromReader", func(t *testing.T) {
		t.Run("when passed a empty reader", func(t *testing.T) {
			emptyReader := strings.NewReader("")
			fileList := ReadFileListFromReader(emptyReader)
			if len(fileList) > 0 {
				t.Errorf("expected filelist to be empty. got: %v: %s", len(fileList), fileList)
			}
		})

		t.Run("when passed a populated reader", func(t *testing.T) {
			control := "file1 file2"
			emptyReader := strings.NewReader(control)
			fileList := ReadFileListFromReader(emptyReader)
			if len(fileList) == 2 {
				t.Errorf("expected filelist to contain given files. expected: %s got: %s", control, fileList)
			}
		})
	})
}
