package main

import  (
	"fmt"
	"net/http"
	"io/ioutil"
	// "io"
	"encoding/json"
	// "strings"
	"log"
	"time"
)

type covid19 struct {
	Number  int `json:"number"`
}

func main() {
	fmt.Println("Starting the application...")
	url := "http://api.open-notify.org/astros.json"
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "spacecount-tutorial")
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	data := covid19{}
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return
	}
	fmt.Println(data.Number)
}
