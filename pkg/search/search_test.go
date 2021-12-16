package search

import (
	"os"
	"testing"
)

func TestSearch_GetColNum(t *testing.T) {
	filepath := "testfile.txt"
	_, err := os.ReadFile(filepath)

	if os.IsNotExist(err) {
		t.Error("file does not exist")
		return
	}
	if err != nil {
		t.Errorf("error ocured while trying to open file %v", err)
	}

	position := GetColNum(filepath, "Hello")
	want := 1
	if position != int64(want) {
		t.Errorf("want: %v, get: %v", want, position)
	}

	position = GetColNum(filepath, "world")
	want = 7
	if position != int64(want) {
		t.Errorf("want: %v, get: %v", want, position)
	}

}

func TestSearch_GetLineNum(t *testing.T) {
	filepath := "testfile.txt"

	line := GetLineNum(filepath, "Hello")
	want := 1
	if line != int64(want) {
		t.Errorf("want: %v, get: %v", want, line)
	}

	line2 := GetLineNum(filepath, "This")
	want2 := 2
	if line2 != int64(want2) {
		t.Errorf("want: %v, get: %v", want2, line2)
	}

}

func TestSearch_GetLine(t *testing.T) {
	filepath := "testfile.txt"
	line := GetLine(filepath, "This")
	want := "This is Go"
	if line != want {
		t.Errorf("want: %v, get: %v", want, line)
	}
}
