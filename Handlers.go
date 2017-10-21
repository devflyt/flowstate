package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/fsouza/go-dockerclient"
)


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "FlowState API")
}

func GetImages(w http.ResponseWriter, r *http.Request) {

	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	images, err := client.ListImages(docker.ListImagesOptions{All: false})
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(images)

}

func GetImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageId := vars["id"]
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	container, err := client.ImageHistory(imageId)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(container)
}

func GetContainers(w http.ResponseWriter, r *http.Request)  {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	containers, err := client.ListContainers(docker.ListContainersOptions{All: false})
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(containers)

}
