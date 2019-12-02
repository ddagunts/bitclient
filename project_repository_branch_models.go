package bitclient

import (
	"fmt"
)

// Those branchmodel/configuration calls are not officially documented. Thanks to this comment for detailing usage:
// https://jira.atlassian.com/browse/BSERV-5411?focusedCommentId=1517096&page=com.atlassian.jira.plugin.system.issuetabpanels%3Acomment-tabpanel#comment-1517096
func (bc *BitClient) GetBranchingModel(projectKey, repositorySlug string) (BranchingModel, error) {
	rError := new(ErrorResponse)

	response := BranchingModel{}

	url := fmt.Sprintf("/rest/branch-utils/1.0/projects/%s/repos/%s/branchmodel/configuration", projectKey, repositorySlug)
	resp, _ := bc.sling.New().Get(url).Receive(&response, rError)

	resp, err := bc.checkReponse(resp, rError)

	return response, err
}

func (bc *BitClient) SetBranchingModel(projectKey, repositorySlug string, settings BranchingModel) error {
	rError := new(ErrorResponse)

	url := fmt.Sprintf("/rest/branch-utils/1.0/projects/%s/repos/%s/branchmodel/configuration", projectKey, repositorySlug)
	resp, _ := bc.sling.New().Put(url).BodyJSON(settings).Receive(nil, rError)

	resp, err := bc.checkReponse(resp, rError)

	return err
}

type SetDefaultBranchRequest struct {
	ID string `json:"id"`
}

func (bc *BitClient) SetDefaultBranch(projectKey, repositorySlug string, params SetDefaultBranchRequest) error {
	_, err := bc.DoPut(
		fmt.Sprintf("/projects/%s/repos/%s/branches/default", projectKey, repositorySlug),
		params,
		nil,
	)

	return err
}
