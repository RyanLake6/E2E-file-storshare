package auth

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

type NextcloudAuth struct {
	BaseURL string
	Client  *http.Client
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

func (auth *NextcloudAuth) Login(username, password string) error {
	req, err := http.NewRequest("GET", auth.BaseURL+"/status.php", nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(username, password)

	resp, err := auth.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Logged in successfully.")
	} else {
		return fmt.Errorf("failed to log in, status code: %d", resp.StatusCode)
	}
	return nil
}
