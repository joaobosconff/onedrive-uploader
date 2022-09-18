package sdk

import (
	"errors"
	"net/http"
	"strings"
)

func (client *Client) Info(path string) (*DriveItem, error) {
	path = strings.TrimSuffix(path, "/")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	url := GraphURL + "me" + client.Config.Root + ":" + path
	if path == "/" {
		url = GraphURL + "me" + client.Config.Root
	}
	status, data, err := client.httpGet(url, nil)
	if err != nil {
		return nil, err
	}
	if status == http.StatusNotFound {
		return nil, errors.New("path not found")
	}
	if status != http.StatusOK {
		return nil, client.handleResponseError(status, data)
	}
	var driveItem DriveItem
	if err := UnmarshalJSON(&driveItem, data); err != nil {
		return nil, err
	}
	return &driveItem, nil
}
