package nextcloud

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type NextcloudFiles struct {
	BaseURL string
	Auth    *http.Client
}

func NewNextcloudFiles(baseURL string, auth *http.Client) *NextcloudFiles {
	return &NextcloudFiles{
		BaseURL: baseURL,
		Auth:    auth,
	}
}


func (files *NextcloudFiles) ListFiles(path string, token string) error {
	req, err := http.NewRequest("PROPFIND", files.BaseURL+"/remote.php/webdav/"+path, nil)
	if err != nil {
		return err
	}

	if token == "" {
		fmt.Println("No set token, please login")
		return fmt.Errorf("no login token set")
	}

	req.Header.Set("Authorization", "Bearer "+ token)
	resp, err := files.Auth.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusMultiStatus {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(body))
	} else {
		return fmt.Errorf("failed to list files, status code: %d", resp.StatusCode)
	}
	return nil
}

func (files *NextcloudFiles) UploadFile(localPath, remotePath string) error {
	file, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer file.Close()

	req, err := http.NewRequest("PUT", files.BaseURL+"/remote.php/webdav/"+remotePath, file)
	if err != nil {
		return err
	}
	resp, err := files.Auth.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		fmt.Println("File uploaded successfully.")
	} else {
		return fmt.Errorf("failed to upload file, status code: %d", resp.StatusCode)
	}
	return nil
}

func (files *NextcloudFiles) DownloadFile(remotePath, localPath string) error {
	req, err := http.NewRequest("GET", files.BaseURL+"/remote.php/webdav/"+remotePath, nil)
	if err != nil {
		return err
	}
	resp, err := files.Auth.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		out, err := os.Create(localPath)
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return err
		}
		fmt.Println("File downloaded successfully.")
	} else {
		return fmt.Errorf("failed to download file, status code: %d", resp.StatusCode)
	}
	return nil
}
