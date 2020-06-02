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

var bjson = "https://a.4cdn.org/" + board + "/threads.json"

type Fourchan []struct {
	Page    int `json:"page"`
	Threads []struct {
		No           int `json:"no"`
		LastModified int `json:"last_modified"`
		Replies      int `json:"replies"`
	} `json:"threads"`
}

type Fourthread struct {
	Posts []struct {
		No            int    `json:"no"`
		Now           string `json:"now"`
		Name          string `json:"name"`
		Sub           string `json:"sub,omitempty"`
		Com           string `json:"com"`
		Filename      string `json:"filename,omitempty"`
		Ext           string `json:"ext,omitempty"`
		W             int    `json:"w,omitempty"`
		H             int    `json:"h,omitempty"`
		TnW           int    `json:"tn_w,omitempty"`
		TnH           int    `json:"tn_h,omitempty"`
		Tim           int    `json:"tim,omitempty"`
		Time          int    `json:"time"`
		Md5           string `json:"md5,omitempty"`
		Fsize         int    `json:"fsize,omitempty"`
		Resto         int    `json:"resto"`
		Bumplimit     int    `json:"bumplimit,omitempty"`
		Imagelimit    int    `json:"imagelimit,omitempty"`
		SemanticURL   string `json:"semantic_url,omitempty"`
		CustomSpoiler int    `json:"custom_spoiler,omitempty"`
		Replies       int    `json:"replies,omitempty"`
		Images        int    `json:"images,omitempty"`
		UniqueIps     int    `json:"unique_ips,omitempty"`
		Spoiler       int    `json:"spoiler,omitempty"`
	} `json:"posts"`
}

func main() {
	r, err := http.Get(bjson)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	var fchan Fourchan
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

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

	var fthread Fourthread
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

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
