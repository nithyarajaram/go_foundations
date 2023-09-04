package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func main() {

	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("Could not open file %v", err)
	}
	defer file.Close()

	wordFrequency(file)

}

var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

func wordFrequency(r io.Reader) (map[string]int, error) {
	scanner := bufio.NewScanner(r)
	lnum := 0
	for scanner.Scan() {
		words := wordRe.FindAllString(scanner.Text(), -1)
		if len(words) != 0 {
			fmt.Println(words)
			lnum++
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	fmt.Println(lnum)
	return nil, nil
}
