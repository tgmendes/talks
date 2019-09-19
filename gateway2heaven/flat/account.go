package account

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"time"

)

// ServiceToken represents an IDKit Service Token payload
type ServiceToken struct {
	AccessTkn string `json:"access_token"`
	ExpiresIn int    `json:"expires_in"`
}

type AccountStatus struct {
	Email    string
	Verified bool
}
type TokenClient interface {
	ServiceToken() (*ServiceToken, error)
}

// Client represents the account client.
type Client struct {
	host       string
	clientID   string
	tknClient  TokenClient
	token      *ServiceToken
	tknCreated time.Time
}

// NewClient creates a custom account client for IdKit account management.
func NewClient(host, clientID string, tknClient TokenClient) *Client {
	cl := Client{
		host:     host,
		clientID: clientID,
		tknClient: tknClient,
	}
	return &cl
}

// VerificationStatus checks the account status of a given client email.
func (cl *Client) VerificationStatus(email string) (*AccountStatus, error) {
	now := time.Now()
	elapsed := int(now.Sub(cl.tknCreated).Seconds())

	// 5 minutes before the token expires it is refreshed
	if cl.token == nil || elapsed >= (cl.token.ExpiresIn-300) {
		tkn, err := cl.tknClient.ServiceToken()
		if err != nil {
			return nil, err
		}
		cl.token = tkn
		cl.tknCreated = time.Now()
	}

	baseURL := cl.host + "/identity/" + email + "/verificationStatus"
	reqBody := []byte(fmt.Sprintf(`{ "appId":"%s", "language":"en" }`, cl.clientID))

	req, err := http.NewRequest(http.MethodPost, baseURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+cl.token.AccessTkn)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 906:
		return &AccountStatus{
			Email:    email,
			Verified: true,
		}, nil
	case 901:
		return &AccountStatus{
			Email:    email,
			Verified: false,
		}, nil
	case http.StatusNotFound:
		return nil, errors.New("account not found")
	}

	return nil, fmt.Errorf("unexpected response from IDKit: status %d", resp.StatusCode)
}
