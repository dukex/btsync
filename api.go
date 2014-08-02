package btsync

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type BTSync struct {
	Endpoint *url.URL
}

type Folder struct {
	Dir      string `json:"dir"`
	Secrect  string `json:"secret"`
	Size     int64  `json:"size"`
	Type     string `json:"type"`
	Files    int64  `json:"files"`
	Error    int    `json:"error"`
	Indexing int    `json:"indexing"`
}

type File struct {
	HavePieces  int    `json:"have_pieces"`
	Name        string `json:"name"`
	Size        int64  `json:"size"`
	State       string `json:"state"`
	TotalPieces int    `json:"total_pieces"`
	Type        string `json:"type"`
	Download    int    `json:"download"`
}

type BTError struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

func New(endpointURL string) (*BTSync, error) {
	endpoint, err := url.Parse(endpointURL)
	if err != nil {
		return nil, err
	}
	endpoint.Path = "api"
	return &BTSync{endpoint}, nil
}

func (b *BTSync) GetFolders(secret string) ([]Folder, error) {
	var folders []Folder
	params := map[string]string{"secret": secret}
	data, err := b.request("get_folders", params)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &folders)
	return folders, err
}

func (b *BTSync) AddFolder(dir, secret, selectiveSync string) (BTError, error) {
	var status BTError
	params := map[string]string{
		"dir":            dir,
		"selective_sync": selectiveSync,
		"secret":         secret,
	}
	data, err := b.request("add_folder", params)
	if err != nil {
		return BTError{}, err
	}
	json.Unmarshal(data, &status)
	return status, err
}

func (b *BTSync) RemoveFolder(secret string) (BTError, error) {
	var status BTError
	params := map[string]string{
		"secret": secret,
	}
	data, err := b.request("remove_folder", params)
	if err != nil {
		return BTError{}, err
	}
	json.Unmarshal(data, &status)
	return status, err
}

func (b *BTSync) GetFiles(secret, path string) ([]File, error) {
	var files []File
	params := map[string]string{"secret": secret, "path": path}
	data, err := b.request("get_files", params)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &files)
	return files, err
}

func (b *BTSync) request(method string, params map[string]string) ([]byte, error) {
	requestURL := b.Endpoint
	q := url.Values{}
	q.Set("method", method)

	for k, v := range params {
		if v != "" {
			q.Set(k, v)
		}
	}
	requestURL.RawQuery = q.Encode()

	log.Println(requestURL.String())
	response, err := http.Get(requestURL.String())
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
