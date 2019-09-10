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

type SSHKey struct {
	Text  string `json:"text"`
	Label string `json:"label"`
}

type AddSSHKeyRequest struct {
	Key        SSHKey `json:"key"`
	Permission string `json:"permission"`
}

func (bc *BitClient) CreateProject(params CreateProjectRequest) (Project, error) {
	response := Project{}

	_, err := bc.DoPost(
		fmt.Sprintf("/projects/"),
		params,
		&response,
	)

	return response, err
}

func (bc *BitClient) DeleteProject(projectKey string) error {

	_, err := bc.DoDeleteUrl(
		fmt.Sprintf("/projects/%s/", projectKey),
		nil,
		nil,
	)

	return err
}

func (bc *BitClient) AddSSHKeyToProject(projectKey string, params AddSSHKeyRequest) error {
	keysBaseUri := "/rest/keys/1.0"
	uri := fmt.Sprintf("/projects/%s/ssh", projectKey)
	rError := new(ErrorResponse)

	resp, _ := bc.sling.New().Post(keysBaseUri+uri).BodyJSON(params).Receive(nil, rError)
	_, err := bc.checkReponse(resp, rError)
	return err
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
