/*
Copyright 2020 CyVerse
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package driver

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// mount client type
type ClientType string

// mount client types
const (
	// IRODSFuseType is for iRODS FUSE
	IRODSFuseType ClientType = "irodsfuse"
	// WebdavType is for WebDav client (Davfs2)
	WebdavType ClientType = "webdav"
)

// IRODSFuseConnection class
type IRODSFuseConnection struct {
	URL        string
	User       string
	Password   string
	ClientUser string // if this field has a value, user and password fields have proxy user info
}

// WebDAVConnection class
type WebDAVConnection struct {
	URL      string
	User     string
	Password string
}

// NewIRODSFuseConnection returns a new instance of IRODSFuseConnection
func NewIRODSFuseConnection(url string, user string, password string, clientUser string) *IRODSFuseConnection {
	return &IRODSFuseConnection{
		URL:        url,
		User:       user,
		Password:   password,
		ClientUser: clientUser,
	}
}

// NewWebDAVConnection returns a new instance of WebDAVConnection
func NewWebDAVConnection(url string, user string, password string) *WebDAVConnection {
	return &WebDAVConnection{
		URL:      url,
		User:     user,
		Password: password,
	}
}

// ExtractClientType extracts Client value from param map
func ExtractClientType(params map[string]string, secrets map[string]string, defaultClient ClientType) ClientType {
	client := ""
	for k, v := range secrets {
		if strings.ToLower(k) == "client" {
			client = v
			break
		}
	}

	for k, v := range params {
		if strings.ToLower(k) == "client" {
			client = v
			break
		}
	}

	return GetValidClientType(client, defaultClient)
}

// IsValidClientType checks if given client string is valid
func IsValidClientType(client string) bool {
	switch client {
	case string(IRODSFuseType):
		return true
	case string(WebdavType):
		return true
	default:
		return false
	}
}

// GetValidClientType checks if given client string is valid
func GetValidClientType(client string, defaultClient ClientType) ClientType {
	switch client {
	case string(IRODSFuseType):
		return IRODSFuseType
	case string(WebdavType):
		return WebdavType
	default:
		return defaultClient
	}
}

// ExtractIRODSFuseConnection extracts IRODSFuseConnection value from param map
func ExtractIRODSFuseConnection(params map[string]string, secrets map[string]string) (*IRODSFuseConnection, error) {
	var user, password, clientUser, url string

	for k, v := range secrets {
		switch strings.ToLower(k) {
		case "user":
			user = v
		case "password":
			password = v
		case "clientuser":
			// for proxy
			clientUser = v
		case "url":
			url = v
		default:
			// ignore
		}
	}

	for k, v := range params {
		switch strings.ToLower(k) {
		case "user":
			user = v
		case "password":
			password = v
		case "clientuser":
			// for proxy
			clientUser = v
		case "url":
			url = v
		default:
			// ignore
		}
	}

	// user and password fields are optional
	// if user is not given, it is regarded as anonymous user
	if len(user) == 0 {
		user = "anonymous"
	}

	// password can be empty for anonymous access
	if len(password) == 0 && user != "anonymous" {
		return nil, status.Error(codes.InvalidArgument, "Argument password is empty")
	}

	if len(url) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Argument url is empty")
	}

	conn := NewIRODSFuseConnection(url, user, password, clientUser)
	return conn, nil
}

// ExtractWebDAVConnection extracts WebDAVConnection value from param map
func ExtractWebDAVConnection(params map[string]string, secrets map[string]string) (*WebDAVConnection, error) {
	var user, password, url string

	for k, v := range secrets {
		switch strings.ToLower(k) {
		case "user":
			user = v
		case "password":
			password = v
		case "url":
			url = v
		default:
			// ignore
		}
	}

	for k, v := range params {
		switch strings.ToLower(k) {
		case "user":
			user = v
		case "password":
			password = v
		case "url":
			url = v
		default:
			// ignore
		}
	}

	// user and password fields are optional
	// if user is not given, it is regarded as anonymous user
	if len(user) == 0 {
		user = "anonymous"
	}

	// password can be empty for anonymous access
	if len(password) == 0 && user != "anonymous" {
		return nil, status.Error(codes.InvalidArgument, "Argument password is empty")
	}

	if len(url) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Argument url is empty")
	}

	conn := NewWebDAVConnection(url, user, password)
	return conn, nil
}
