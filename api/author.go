package api

import (
	"net/http"
	"time"
)

type AuthorRequest struct {
	username string
}

func NewAuthorRequest(username string) *AuthorRequest {
	return &AuthorRequest{username: username}
}

// https://api.gitbook.com/author/{user_name}
func (r *AuthorRequest) Uri() string {
	return "author/" + r.username
}

func (r *AuthorRequest) Method() string {
	return http.MethodGet
}

func (r *AuthorRequest) Response() Responseer {
	return new(Author)
}

type Author struct {
	CommonResponse
	ID          string `json:"id"`
	HasMigrated bool   `json:"hasMigrated"`
	Type        string `json:"type"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	Website     string `json:"website"`
	Bio         string `json:"bio"`
	Verified    bool   `json:"verified"`
	Locked      bool   `json:"locked"`
	SiteAdmin   bool   `json:"site_admin"`
	Urls        struct {
		Profile string `json:"profile"`
		Stars   string `json:"stars"`
		Avatar  string `json:"avatar"`
	} `json:"urls"`
	Permissions struct {
		Edit  string `json:"edit"`
		Admin string `json:"admin"`
	} `json:"permissions"`
	Dates struct {
		Created *time.Time
	} `json:"dates"`
}
