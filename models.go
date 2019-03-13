package main

type Project struct {
	ID                string `json:"projectId"`
	Description       int    `json:"description"`
	Name              string `json:"name"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
}
