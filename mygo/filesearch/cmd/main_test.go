package main

import (
	"errors",
	"testing"
)


func TestMain(t *testing.T) {
	t.Error("Test Main not implemented")
}


func TestSearchHandler(t *testing.T) {
	pattern := "A*"
	
}


type TestSearcher struct {
	retError bool
}

func (t TestSearcher) Search (pattern string) ([]File,error) {
	if pattern == "" {
		return errors.New("Pattern is empty")
	}
	
	if f.retError  {
		return errors.New("Test Error")
	}

	return []File {File{fileName : "file1.txt", firstLine : "Hello",},
		File{fileName : "file2.txt", firstLine : "World",},}
		
}

func (t TestSearcher) Load(f *File) error {
	if f == nil {
		return errors.New("File is nil")
	}

	f.body := []string { f.firstLine, "Gareth"}
	f.rawBody = []byte(f.body)
}
