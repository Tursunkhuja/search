package search

import (
	"context"
	"testing"
)

func TestAll_success(t *testing.T) {
	ch := All(context.Background(), "Hello", []string{"testfile.txt"})

	_, err := <-ch

	if !err {
		t.Error(err)
	}
}

func TestAny_success(t *testing.T) {
	ch := Any(context.Background(), "Hello", []string{"testfile.txt"})

	_, err := <-ch

	if !err {
		t.Error(err)
	}
}
