package nextcloud

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

// Define the structure that matches the XML response
type Multistatus struct {
	Responses []Response `xml:"response"`
}

type Response struct {
	Href     string    `xml:"href"`
	Propstat Propstat  `xml:"propstat"`
}

type Propstat struct {
	Prop   Prop   `xml:"prop"`
	Status string `xml:"status"`
}

type Prop struct {
	LastModified        string `xml:"getlastmodified"`
	ResourceType        string `xml:"resourcetype>collection"`
	QuotaUsedBytes      string `xml:"quota-used-bytes"`
	QuotaAvailableBytes string `xml:"quota-available-bytes"`
	Etag                string `xml:"getetag"`
	ContentLength       string `xml:"getcontentlength"`
	ContentType         string `xml:"getcontenttype"`
}


func (files *NextcloudFiles) ListFiles(path string, token string, allDetails bool) error {
	req, err := http.NewRequest("PROPFIND", files.BaseURL+"/remote.php/webdav/"+path, nil)
	if err != nil {
		return err
	}

	if token == "" {
		fmt.Println("No set token, please login")
		return fmt.Errorf("no login token set")
	}

	req.Header.Set("Authorization", "Bearer "+token)
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

		// Removing namespace prefixes for easier processing
		bodyStr := strings.ReplaceAll(string(body), "d:", "")
		bodyStr = strings.ReplaceAll(bodyStr, "oc:", "")
		bodyStr = strings.ReplaceAll(bodyStr, "s:", "")
		bodyStr = strings.ReplaceAll(bodyStr, "nc:", "")

		var multistatus Multistatus
		err = xml.Unmarshal([]byte(bodyStr), &multistatus)
		if err != nil {
			return err
		}

		// Printing the file information
		for _, response := range multistatus.Responses {
			fmt.Printf("Href: %s\n", response.Href)
			fmt.Printf("Last Modified: %s\n", response.Propstat.Prop.LastModified)

			if allDetails {
				fmt.Printf("Resource Type: %s\n", response.Propstat.Prop.ResourceType)
				fmt.Printf("Quota Used Bytes: %s\n", response.Propstat.Prop.QuotaUsedBytes)
				fmt.Printf("Quota Available Bytes: %s\n", response.Propstat.Prop.QuotaAvailableBytes)
				fmt.Printf("ETag: %s\n", response.Propstat.Prop.Etag)
				fmt.Printf("Content Length: %s\n", response.Propstat.Prop.ContentLength)
				fmt.Printf("Content Type: %s\n", response.Propstat.Prop.ContentType)
			}

			fmt.Println("---")
		}

	} else {
		return fmt.Errorf("failed to list files, status code: %d", resp.StatusCode)
	}
	return nil
}

func (files *NextcloudFiles) UploadFile(localPath, remotePath, token string) error {
	file, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer file.Close()

	req, err := http.NewRequest("PUT", files.BaseURL+"/remote.php/webdav/"+remotePath, file)
	if err != nil {
		return err
	}

	// Set Content-Type header (adjust as needed)
	req.Header.Set("Content-Type", "application/octet-stream")

	if token == "" {
		fmt.Println("No set token, please login")
		return fmt.Errorf("no login token set")
	}

	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := files.Auth.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		fmt.Println("File uploaded successfully.")
	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Failed to read response body: %v\n", err)
		}
		bodyString := string(bodyBytes)
		return fmt.Errorf("failed to upload file, status code: %d, body: %s", resp.StatusCode, bodyString)
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
