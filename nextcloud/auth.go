package nextcloud

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

type NextcloudAuth struct {
	BaseURL string
	Client  *http.Client
}

type LoginResponse struct {
	Poll struct {
		Token    string `json:"token"`
		Endpoint string `json:"endpoint"`
	} `json:"poll"`
	Login string `json:"login"`
}

type PollResponse struct {
	Server      string `json:"server"`
	LoginName   string `json:"loginName"`
	AppPassword string `json:"appPassword"`
}

func NewNextcloudAuth(baseURL string) *NextcloudAuth {
	// disabled certificate validation in the HTTP client due to a certificate signed by unkown authority bug
	// ** This is done in a secure and trusted environment, all users must be aware that certificate validation is disabled **
	return &NextcloudAuth{
		BaseURL: baseURL,
		Client:  &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},},
	}
}


func (auth *NextcloudAuth) Login() (*PollResponse, error) {
	// Step 1: Initiate the login flow
	loginResp, err := auth.initiateLogin()
	if err != nil {
		return nil, err
	}

	// Step 2: Open the login URL in the default browser
	err = openBrowser(loginResp.Login)
	if err != nil {
		return nil, err
	}

	// Step 3: Poll the endpoint for the authentication token
	return auth.pollForToken(loginResp.Poll.Endpoint, loginResp.Poll.Token)
}

func (auth *NextcloudAuth) initiateLogin() (*LoginResponse, error) {
	req, err := http.NewRequest("POST", auth.BaseURL+"/index.php/login/v2", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := auth.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var loginResp LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return nil, err
	}
	return &loginResp, nil
}

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch os := runtime.GOOS; os {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default:
		cmd = "xdg-open"
		args = []string{url}
	}
	return exec.Command(cmd, args...).Start()
}

func (auth *NextcloudAuth) pollForToken(endpoint, token string) (*PollResponse, error) {
	fmt.Print("Awaiting browser login")
	for {
		data := []byte(fmt.Sprintf("token=%s", token))
		req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		client := &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},}
		fmt.Print(".")
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode == http.StatusOK {
			var pollResp PollResponse
			if err := json.NewDecoder(resp.Body).Decode(&pollResp); err != nil {
				return nil, err
			}
			defer resp.Body.Close()
			return &pollResp, nil
		} else if resp.StatusCode != http.StatusNotFound {
			defer resp.Body.Close()
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		time.Sleep(1 * time.Second)
	}
}