package agent

import (
	"encoding/json"
	"errors"
	"ex-server/internal/task/entity"
	"fmt"
	"net/http"
)

var ErrForbidden error = errors.New("forbidden")

func Init(url string) *AuthAgent {
	return &AuthAgent{url: url}
}

type AuthAgent struct {
	url string
}

func (agent *AuthAgent) Check(token string) (*entity.User, error) {
	client := http.Client{}
	req, err := http.NewRequest(
		"GET", fmt.Sprintf("%s/check", agent.url), nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Token", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, ErrForbidden
	}

	var user entity.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, ErrForbidden
	}

	return &user, nil
}
