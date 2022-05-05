package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Entries []Entries
	Cursor  string
	HasMore bool
}

type Entries struct {
	Tag  string `json:".tag"`
	Name string `json:"name"`
	Path string `json:"path_lower"`
	Id   string `json:"id"`
}

var token = "Bearer sl.BHCO8jgBhjcpXkgTjzDGIr50l43YHZZndyfklZmZYIwah5wdSwRdH3jrJ_8kyFfKiP98Gc-yJf9lt3bwNtKpd_KZLmiluHgv3P6rr8QG-LI9dRXWfG40NMDcSR4YOLWde9YEZCAxXIGo"

func deleteFile(path string) {
	url := "https://api.dropboxapi.com/2/files/delete_v2"

	var jsonStr = []byte(`{
    "path": "` + path + `"}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Add("Authorization", token)

	req.Header.Set("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

}

func getData() {
	url := "https://api.dropboxapi.com/2/files/list_folder"

	var jsonStr = []byte(`{"include_deleted": false,
    "include_has_explicit_shared_members": false,
    "include_media_info": false,
    "include_mounted_folders": true,
    "include_non_downloadable_files": true,
    "path": "/test",
    "recursive": true}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Add("Authorization", token)
	//request.Header.Add("cursor", response.Cursor)
	req.Header.Set("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response Response
	json.Unmarshal(body, &response)

	//fmt.Println(response.Entries)
	for e, f := range response.Entries {
		fmt.Println("File numero:", e)
		fmt.Println("Tag:", f.Tag)
		fmt.Println("Name:", f.Name)
		if string(f.Name[0]) == "~" {
			deleteFile(f.Path)
			fmt.Println("Path:", f.Path)
		}
		fmt.Println("Path:", f.Path)
		fmt.Println("Id:", f.Id)
		fmt.Println("--------------------------------")
	}

}

func main() {
	getData()
}
