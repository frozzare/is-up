package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Response struct {
	Status_code int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%s %s", getColorString(9, 3, "Error:"), getColorString(9, 0, "Missing url"))
		return
	}

	if isUp(os.Args[1]) {
		fmt.Println(getColorString(9, 2, "Up"))
	} else {
		fmt.Println(getColorString(9, 3, "Down"))
	}
}

// Get string with terminal color.
func getColorString(firstColor int, lastColor int, message string) string {
	return fmt.Sprintf("\u001b[%d%dm%s\u001b[0m", firstColor, lastColor, message)
}

// Check if url is up or down.
func isUp(url string) bool {
	r, _ := regexp.Compile(`^https?:\/\/`)

	url = r.ReplaceAllString(url, "")
	url = "http://isitup.org/" + url + ".json"

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var data Response

	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal(err)
	}

	if data.Status_code == 3 {
		fmt.Printf("%s %s", getColorString(9, 3, "Error:"), getColorString(9, 0, "Invalid domain"))
		println()
		os.Exit(0)
	}

	return data.Status_code == 1
}
