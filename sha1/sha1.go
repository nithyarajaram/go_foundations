package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	signature, err := sha("http.log.gz")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	fmt.Println("Sha1 Signature:", signature)

	signature, err = sha("sha1.go")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	fmt.Println("Sha1 Signature:", signature)
}

func sha(filename string) (string, error) {
	//Read file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Cannot open file: %s", err)
	}
	defer file.Close()
	var r io.Reader = file

	if strings.HasSuffix(filename, ".gz") {

		//Unzip the file

		gr, err := gzip.NewReader(file)
		if err != nil {
			return "", err
		}
		defer gr.Close()
		r = gr

	}

	//shasum

	h := sha1.New()
	if _, err = io.Copy(h, r); err != nil {
		return " ", err
	}

	signature := h.Sum(nil)
	return fmt.Sprintf("%x", signature), nil

}
