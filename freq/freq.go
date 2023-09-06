package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {

	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("Could not open file %v", err)
	}
	defer file.Close()

	//fmt.Println(wordFrequency(file))
	//fmt.Println(maxWord(frequency))
	fmt.Println(mostCommon(file))
}

var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

func maxWord(frequency map[string]int) (string, int, error) {

	if len(frequency) == 0 {
		return "", 0, fmt.Errorf("empty map")
	}
	maxW, maxN := "", 0
	for word, count := range frequency {
		if count > maxN {
			maxW, maxN = word, count
		}
	}
	return maxW, maxN, nil
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

func mapDemo() {
	var stocks map[string]float64 //symbol -> price
	sym := "znga"
	//price := stocks[sym]
	//fmt.Printf("%s -> %.2f\n", sym, price)
	stocks = map[string]float64{
		sym:    132.5,
		"APPL": 239.2,
	}

	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> %.2f\n", sym, price)
	} else {
		fmt.Printf("%s not found", sym)
	}

	for k, v := range stocks {
		fmt.Println(k, v)
	}
	for _, v := range stocks {
		fmt.Println(v)
	}
}

func mostCommon(r io.Reader) (string, int, error) {
	frequency, err := wordFrequency(r)
	if err != nil {
		return "", 0, err
	}
	return maxWord(frequency)
}
