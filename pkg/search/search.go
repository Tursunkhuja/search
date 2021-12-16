package search

import (
	"context"
	"os"
	"strings"
	"sync"
)

//Resuлt описывает один результат поиска.
type Result struct {
	//The phrase you were looking for
	Phrase string
	//Entire line in which the entry was found without \ n or \ r \ n at the end
	Line string
	//line number (starting from 1) where the entry was found
	LineNum int64
	//the position number (starting from 1) at which the entry was found
	ColNum int64
}

//All ищет все вхождение phrase в текстовых файлах files
func All(ctx context.Context, phrase string, files []string) <-chan []Result {

	ch := make(chan []Result)

	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(ctx)

	for i := 0; i < len(files); i++ {
		wg.Add(1)

		go func(ctx context.Context, filename string, i int, ch chan<- []Result) {
			defer wg.Done()
			res := FindAllMatchTextInFile(phrase, filename)

			if len(res) > 0 {
				ch <- res
			}

		}(ctx, files[i], i, ch)

		go func() {
			defer close(ch)
			wg.Wait()
		}()

	}
	cancel()
	return ch
}

func FindAllMatchTextInFile(phrase, fileName string) (res []Result) {
	read, _ := os.ReadFile(fileName)
	fstr := string(read)
	filestr := strings.Split(fstr, "\n")

	if len(filestr) > 0 {

		filestr = filestr[:len(filestr)-1]
	}
	for i, line := range filestr {

		if strings.Contains(line, phrase) {

			result := Result{
				Phrase:  phrase,
				Line:    line,
				LineNum: int64(i + 1),
				ColNum:  int64(strings.Index(line, phrase)) + 1,
			}
			res = append(res, result)
		}
	}
	return res
}
