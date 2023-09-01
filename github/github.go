package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Reply struct {
	Name         string
	Public_Repos int
}

func main() {

	fmt.Println(githubInfo("tebeka"))

}

func githubInfo(login string) (string, int, error) {

	resp, err := http.Get("https://api.github.com/users/" + login)
	if err != nil {
		log.Fatalf("Error %s:", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error %s:", resp.Status)
	}
	// fmt.Printf("Content-Type %s", resp.Header.Get("Content-Type"))

	/* if _, err := io.Copy(os.Stdout, resp.Body); err != nil {

		log.Fatalf("Error: Can't copy %s", err)
	} */

	var r Reply

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("Error: Can't decode %s", err)
	}
	return r.Name, r.Public_Repos, nil

}
