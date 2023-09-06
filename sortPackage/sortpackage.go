package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {

	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("Could not open file %v", err)
	}
	defer file.Close()

	fmt.Println(mostCommonNWords(file, 20))

}

var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

func mostCommonNWords(r io.Reader, n int) ([]string, error) {

	frequency, err := wordFrequency(r)
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(frequency))
	for key := range frequency {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return frequency[keys[i]] > frequency[keys[j]]
	})
	return keys[:n], nil
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	scanner := bufio.NewScanner(r)
	frequency := make(map[string]int) //word -> count
	lnum := 0
	for scanner.Scan() {
		words := wordRe.FindAllString(scanner.Text(), -1)
		for _, i := range words {
			frequency[strings.ToLower(i)]++
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	fmt.Println(lnum)
	return frequency, nil

}
