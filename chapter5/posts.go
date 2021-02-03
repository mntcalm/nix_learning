package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
)


func main() {
  resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
     if err != nil {
	log.Fatalln(err)
     }
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
	log.Fatalln(err)
    }

  fmt.Println(string(body))
}