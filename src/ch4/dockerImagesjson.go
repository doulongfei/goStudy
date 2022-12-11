package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Images struct {
	Repositories []string `json:"repositories"`
}

type Tag struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func getImagesData(url string, images *Images) Images {
	resp, err := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println("访问异常", err)
	}
	body, _ := io.ReadAll(resp.Body)
	//err = json.Unmarshal(body, &re)
	if err = json.Unmarshal(body, &images); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return Images{}
	}
	return *images
}
func getTagsData(url, image string, tag *Tag) Tag {
	resp, err := http.Get(url + "/v2/" + image + "/tags/list")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println("访问异常", err)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	//err = json.Unmarshal(body, &re)
	if err = json.Unmarshal(body, &tag); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return Tag{}
	}
	return *tag
}

func main() {

	strings := os.Args[1:]
	//fmt.Println(strings)
	var result string
	url := strings[0]
	//fmt.Println(url)
	var images Images
	getImagesData(url+"/v2/_catalog", &images)
	//fmt.Println(images)
	for _, repository := range images.Repositories {
		var tags Tag
		getTagsData(url, repository, &tags)
		//fmt.Println(tags)
		for _, tagOne := range tags.Tags {
			res := url + ":" + tagOne + "|"
			result += res
		}
	}
	fmt.Println(result[:len(result)-1])
}
