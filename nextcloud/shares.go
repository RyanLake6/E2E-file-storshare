package nextcloud

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ShareType int

const (
	ShareTypeUser  ShareType = 0
	ShareTypeGroup ShareType = 1
	ShareTypeLink  ShareType = 3
)

type NextcloudShare struct {
	BaseURL string
	Client  *http.Client
}

type ShareResponse struct {
	XMLName xml.Name `xml:"ocs"`
	Meta    Meta     `xml:"meta"`
	Data    Data     `xml:"data"`
}

type Meta struct {
	Status     string `xml:"status"`
	StatusCode int    `xml:"statuscode"`
	Message    string `xml:"message"`
}

type Data struct {
	ID                    int    `xml:"id"`
	ShareType             int    `xml:"share_type"`
	UIDOwner              string `xml:"uid_owner"`
	DisplayNameOwner      string `xml:"displayname_owner"`
	Permissions           int    `xml:"permissions"`
	CanEdit               int    `xml:"can_edit"`
	CanDelete             int    `xml:"can_delete"`
	STime                 int64  `xml:"stime"`
	Parent                string `xml:"parent"`
	Expiration            string `xml:"expiration"`
	Token                 string `xml:"token"`
	UIDFileOwner          string `xml:"uid_file_owner"`
	Note                  string `xml:"note"`
	Label                 string `xml:"label"`
	DisplayNameFileOwner  string `xml:"displayname_file_owner"`
	Path                  string `xml:"path"`
	ItemType              string `xml:"item_type"`
	ItemPermissions       int    `xml:"item_permissions"`
	MimeType              string `xml:"mimetype"`
	HasPreview            string `xml:"has_preview"`
	StorageID             string `xml:"storage_id"`
	Storage               int    `xml:"storage"`
	ItemSource            int    `xml:"item_source"`
	FileSource            int    `xml:"file_source"`
	FileParent            int    `xml:"file_parent"`
	FileTarget            string `xml:"file_target"`
	ItemSize              int64  `xml:"item_size"`
	ItemMTime             int64  `xml:"item_mtime"`
	ShareWith             string `xml:"share_with"`
	ShareWithDisplayName  string `xml:"share_with_displayname"`
	Password              string `xml:"password"`
	SendPasswordByTalk    string `xml:"send_password_by_talk"`
	URL                   string `xml:"url"`
	MailSend              int    `xml:"mail_send"`
	HideDownload          int    `xml:"hide_download"`
	Attributes            string `xml:"attributes"`
}

func NewNextcloudShare(baseURL string, client *http.Client) *NextcloudShare {
	return &NextcloudShare{
		BaseURL: baseURL,
		Client:  client,
	}
}

func (share *NextcloudShare) CreateShare(remotePath string, shareType ShareType, permissions int, token string, debug bool) (string, error) {
	data := fmt.Sprintf("path=%s&shareType=%d&permissions=%d", remotePath, shareType, permissions)
	req, err := http.NewRequest("POST", share.BaseURL+"/ocs/v2.php/apps/files_sharing/api/v1/shares", strings.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("OCS-APIRequest", "true")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if token == "" {
		fmt.Println("No set token, please login")
		return "", fmt.Errorf("no login token set")
	}

	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := share.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to create share, status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var shareResponse ShareResponse
	err = xml.Unmarshal([]byte(string(body)), &shareResponse)
	if err != nil {
		return "", err
	}

	// Printing extra information if debug flag is included
	if debug {
		fmt.Println("share response headers: ", shareResponse)
		fmt.Println("full response: ", string(body))
		fmt.Println("Token: ", token)
		fmt.Println("Response status code: ", resp.StatusCode)
		fmt.Println("share data to be utilized: ", data)
	}
	
	return shareResponse.Data.URL, nil
}

func (share *NextcloudShare) ListShares(token string) error {
	req, err := http.NewRequest("GET", share.BaseURL+"/ocs/v2.php/apps/files_sharing/api/v1/shares", nil)
	if err != nil {
		return err
	}
	req.Header.Set("OCS-APIRequest", "true")

	if token == "" {
		fmt.Println("No set token, please login")
		return fmt.Errorf("no login token set")
	}

	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := share.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to list shares, status code: %d", resp.StatusCode)
	}

	var body bytes.Buffer
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(body.String())

	return nil
}

func (share *NextcloudShare) DeleteShare(shareID string, token string) error {
	req, err := http.NewRequest("DELETE", share.BaseURL+"/ocs/v2.php/apps/files_sharing/api/v1/shares/"+shareID, nil)
	if err != nil {
		return err
	}
	req.Header.Set("OCS-APIRequest", "true")

	if token == "" {
		fmt.Println("No set token, please login")
		return fmt.Errorf("no login token set")
	}

	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := share.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete share, status code: %d", resp.StatusCode)
	}

	return nil
}
