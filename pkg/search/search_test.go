package search

import (
	"context"
	"log"
	"testing"
)

func TestAll(t *testing.T) {

	ctx := context.Background()
	files := []string{"testfile.txt"}

	ch := All(ctx, "Hello", files)

	s, ok := <-ch

	if !ok {
		t.Errorf("fuction All error +> %v", ok)
	}

	log.Println("---------", s)
}
