package main

import (
    "fmt"
    "net/http"
    "log"
    "io/ioutil"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
/*    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:]) */
        resp, err := http.Get("https://thevirustracker.com/free-api?global=stats")
        if err != nil {
                log.Fatal(err)
        }
        if resp.StatusCode == http.StatusOK {
                bodyBytes, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                        log.Fatal(err)
                }
                bodyString := string(bodyBytes)
                fmt.Fprintf(w, "%s", bodyString)
        }
        defer resp.Body.Close()
}
