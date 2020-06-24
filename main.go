package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var query = strings.Join(os.Args[1:], " ")
var squery = strings.Split(query, "/")
var board = squery[3]
var flongname = squery[4]
var fname2 = strings.Split(flongname, ".")
var fname = fname2[0]

type Fourchan []struct {
	Page    int `json:"page"`
	Threads []struct {
		No int `json:"no"`
	} `json:"threads"`
}

type Fourthread struct {
	Posts []struct {
		No  int `json:"no"`
		Tim int `json:"tim,omitempty"`
	} `json:"posts"`
}

func main() {
	var bjson = "https://a.4cdn.org/" + board + "/threads.json"
	r, err := http.Get(bjson)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var fchan Fourchan
	err = json.Unmarshal(body, &fchan)

	for _, c := range fchan {
		for _, t := range c.Threads {
			scant(t.No)
		}
	}
}

func scant(thread int) {
	time.Sleep(100 * time.Millisecond)
	sure := strconv.Itoa(thread)
	var tjson = "https://a.4cdn.org/" + board + "/thread/" + sure + ".json"

	r, err := http.Get(tjson)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var fthread Fourthread
	err = json.Unmarshal(body, &fthread)

	for _, t := range fthread.Posts {
		tmpname, _ := strconv.Atoi(fname)
		if t.Tim == tmpname {
			tmptno := strconv.Itoa(t.No)
			fmt.Println("https://boards.4chan.org/" + board + "/thread/" + sure + "#" + tmptno)
			os.Exit(0)
		}
	}
}
