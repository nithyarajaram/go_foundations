package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

	//Anonymous struct for one-off usage consuming API's
	var r struct {
		Name      string
		Num_Repos int `json:"public_repos"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}
	return r.Name, r.Num_Repos, nil

}
