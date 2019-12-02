package bitclient

import (
	"fmt"
)

type UpdateRepositoryRequest struct {
	Name     string  `json:"name,omitempty"`
	Forkable bool    `json:"forkable,omitempty"`
	Project  Project `json:"project,omitempty"`
	Public   bool    `json:"public,omitempty"`
}

func (bc *BitClient) UpdateRepository(projectKey string, repositorySlug string, params UpdateRepositoryRequest) (Repository, error) {

	response := Repository{}

	url := fmt.Sprintf("/projects/%s/repos/%s", projectKey, repositorySlug)
	_, err := bc.DoPut(url, params, response)

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

func (bc *BitClient) DeleteRepository(projectKey, repoSlug string) error {

	_, err := bc.DoDeleteUrl(
		fmt.Sprintf("/projects/%s/repos/%s", projectKey, repoSlug),
		nil,
		nil,
	)

	return err
}
