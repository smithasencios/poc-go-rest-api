package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/gitlab/projects/{projectId}", GetGitLabProject).Methods("GET")
	router.HandleFunc("/circleci/projects", GetCircleCIProjects).Methods("GET")
	return router
}
