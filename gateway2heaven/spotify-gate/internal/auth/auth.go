package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Authenticator is used to generate and validate tokens.
type Authenticator struct {
	httpClient   *http.Client
	TokenURL     string
	ClientID     string
	ClientSecret string
	Token        *Token
}

// GenerateToken will do an HTTP call to the token endpoint and generate a token to be used with an API.
func (a *Authenticator) GenerateToken() error {
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", a.ClientID, a.ClientSecret)))

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	r, err := http.NewRequest(http.MethodPost, a.TokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	r.Header.Add("Authorization", fmt.Sprintf("Basic %s", auth))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := a.httpClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tkn := Token{}
	if err = json.NewDecoder(resp.Body).Decode(&tkn); err != nil {
		return err
	}

	a.Token = &tkn
	return nil
}
