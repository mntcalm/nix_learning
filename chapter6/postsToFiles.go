package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	//	"os"
	"strconv"
)

type PostIt struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	TTT := make(chan string)
	RRR := make(chan string)
	NNN := make(chan int)
	var urlN, ell, filePath string
	var pNom int
	//    numPost := make(chan string)
	//    bodyPost := make(chan string)
	for i := 1; i <= 100; i++ {
		urlN = "https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i)
		go oneposter(TTT, RRR, NNN)
		TTT <- urlN
	}

	for i := 1; i <= 100; i++ {
		ell = <-RRR
		pNom = <-NNN
		if ell != "none" {

			filePath = "storage/posts/" + strconv.Itoa(pNom) + ".txt"
			//post, err := os.Create(filePath)
			err := ioutil.WriteFile(filePath, []byte(ell), 0644)
			if err != nil {
				//	post.WriteString(ell)
				fmt.Println("Не могу создать файл!")
			}

		}
	}

}

func oneposter(TTT chan string, RRR chan string, NNN chan int) int {
	ell := ""
	url := <-TTT
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	salo := PostIt{}
	err = json.Unmarshal(body, &salo)
	if err != nil {
		fmt.Println(err)

	}
	ell = "{UserId:" + strconv.Itoa(salo.UserId) + " Id:" + strconv.Itoa(salo.Id) + " Title:" + salo.Title + " Body:" + salo.Body + "}"
	RRR <- ell
	NNN <- salo.Id
	return 0
}
