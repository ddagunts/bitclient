package bitclient

import (
	"fmt"
)

type GetProjectsResponse struct {
	PagedResponse
	Values []Project
}
type CreateProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Key         string `json:"key"`
}

func (bc *BitClient) CreateProjects(params CreateProjectRequest) (Project, error) {
	response := Project{}

	_, err := bc.DoPost(
		fmt.Sprintf("/projects/"),
		params,
		&response,
	)

	return response, err
}

func (bc *BitClient) GetProjects(params PagedRequest) (GetProjectsResponse, error) {

	response := GetProjectsResponse{}

	_, err := bc.DoGet(
		"/projects",
		params,
		&response,
	)

	return response, err
}

type GetRepositoriesResponse struct {
	PagedResponse
	Values []Repository
}

func (bc *BitClient) GetRepositories(projectKey string, params PagedRequest) (GetRepositoriesResponse, error) {

	response := GetRepositoriesResponse{}

	_, err := bc.DoGet(
		fmt.Sprintf("/projects/%s/repos", projectKey),
		params,
		&response,
	)

	return response, err
}

type CreateRepositoryRequest struct {
	Name     string `json:"name"`
	ScmId    string `json:"scmId"`
	Forkable bool   `json:"forkable"`
}

func (bc *BitClient) CreateRepository(projectKey string, params CreateRepositoryRequest) (Repository, error) {

	response := Repository{}

	_, err := bc.DoPost(
		fmt.Sprintf("/projects/%s/repos", projectKey),
		params,
		&response,
	)

	return response, err
}
