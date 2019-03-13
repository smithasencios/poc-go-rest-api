package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const gitLabAccessTokenKey = "Private-Token"
const gitLabAccessTokenValue = "GITLAB_PRIVATE_TOKEN_GOES_HERE"
const circleCIAPIToken = "CIRCLECI_PRIVATE_TOKEN_GOES_HERE"

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func GetCircleCIProjects(w http.ResponseWriter, r *http.Request) {
	url := "https://circleci.com/api/v1.1/projects?circle-token=" + circleCIAPIToken

	// Create new request for Gitlab
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	// Set headers
	req.Header.Set("Cache-Control", "no-cache")

	// Set client timeout
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Send request to Gitlab
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	// Reading Gitlab's response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	respondJSON(w, http.StatusOK, body)
}

func GetGitLabProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectID := params["projectId"]
	url := "http://gitlab.com/api/v4/projects/" + projectID

	// Create new request for Gitlab
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	// Set headers
	req.Header.Set(gitLabAccessTokenKey, gitLabAccessTokenValue)
	req.Header.Set("Cache-Control", "no-cache")

	// Set client timeout
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Send request to Gitlab
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	// Reading Gitlab's response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	respondJSON(w, http.StatusOK, body)
}

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(payload)
}

// // respondError makes the error response with payload as json format
// func respondError(w http.ResponseWriter, code int, message string) {
// 	respondJSON(w, code, map[string]string{"error": message})
// }
