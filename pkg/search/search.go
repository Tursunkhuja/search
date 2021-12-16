package search

import (
	"context"
	"io/ioutil"
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

// FindMatchesInFile finds all phrase occurrences
func FindMatchesInFile(phrase, file string, findingAll bool) ([]Result, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var result []Result = nil
	for i, line := range strings.Split(string(data), "\n") {
		if strings.Contains(line, phrase) {
			found := Result{
				Phrase:  phrase,
				Line:    line,
				LineNum: int64(i + 1),
				ColNum:  int64(strings.Index(line, phrase) + 1),
			}

			result = append(result, found)

			if !findingAll {
				return result, nil
			}
		}
	}

	return result, nil
}

// All is the main function for finding occurrences of phrase in given list of files
func All(ctx context.Context, phrase string, files []string) <-chan []Result {
	ch := make(chan []Result)
	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(ctx)

	for i, file := range files {
		wg.Add(1)

		go func(ctx context.Context, filename string, i int, ch chan<- []Result) {
			defer wg.Done()

			result, _ := FindMatchesInFile(phrase, filename, true)

			if len(result) > 0 {
				ch <- result
			}
		}(ctx, file, i, ch)
	}

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	cancel()
	return ch
}

// Any is the main function for finding one of the occurrences of phrase in given list of files
func Any(ctx context.Context, phrase string, files []string) <-chan Result {
	ch := make(chan Result)
	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(ctx)

	var result []Result
	for _, file := range files {
		current, err := FindMatchesInFile(phrase, file, false)
		if err != nil {
			continue
		}

		if len(current) > 0 {
			result = current

			break
		}
	}

	wg.Add(1)

	go func(ctx context.Context, ch chan<- Result) {
		defer wg.Done()

		if len(result) > 0 {
			ch <- result[0]
		}
		cancel()
	}(ctx, ch)

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	cancel()

	return ch
}
