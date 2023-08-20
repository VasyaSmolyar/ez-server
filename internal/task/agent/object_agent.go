package agent

import (
	"errors"
	"fmt"
	"net/http"
)

var ErrFileNotFound error = errors.New("file was not found")

func NewObjectAgent(url string) *ObjectAgent {
	return &ObjectAgent{url: url}
}

type ObjectAgent struct {
	url string
}

func (agent *ObjectAgent) Check(objectName string) error {
	client := http.Client{}
	req, err := http.NewRequest(
		"GET", fmt.Sprintf("%s/check/%s", agent.url, objectName), nil,
	)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return ErrFileNotFound
	}

	return nil
}
