package main

import (
	"io/ioutil"
	"encoding/json"
)

type video struct{
	Id string
	Title string
	Description string
	Imageurl string
	Url string
}

func getVideos()(videos []video){
	fileBytes, err := ioutil.ReadFile("./videos.json")  
	
	if err != nil {
		panic(err)
	}
	
	err = json.Unmarshal(fileBytes, &videos)
	
	if err != nil {
		panic(err)
	}
	
	return videos
}