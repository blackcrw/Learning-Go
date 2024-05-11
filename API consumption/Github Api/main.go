package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type RequestJsonAPI struct {
	Login             string      `json:"login"`
	ID                int         `json:"id"`
	NodeID            string      `json:"node_id"`
	AvatarURL         string      `json:"avatar_url"`
	GravatarID        string      `json:"gravatar_id"`
	URL               string      `json:"url"`
	HTMLURL           string      `json:"html_url"`
	FollowersURL      string      `json:"followers_url"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	OrganizationsURL  string      `json:"organizations_url"`
	ReposURL          string      `json:"repos_url"`
	EventsURL         string      `json:"events_url"`
	ReceivedEventsURL string      `json:"received_events_url"`
	Type              string      `json:"type"`
	SiteAdmin         bool        `json:"site_admin"`
	Name              string      `json:"name"`
	Company           string      `json:"company"`
	Blog              string      `json:"blog"`
	Location          string      `json:"location"`
	Email             interface{} `json:"email"`
	Hireable          interface{} `json:"hireable"`
	Bio               string      `json:"bio"`
	TwitterUsername   interface{} `json:"twitter_username"`
	PublicRepos       int         `json:"public_repos"`
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

func main() {
	req, err := http.Get("https://api.github.com/users/blackcrw")
	defer req.Body.Close()

	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		return
	}

	var jsn RequestJsonAPI

	if err := json.NewDecoder(req.Body).Decode(&jsn); err != nil {
		log.Println(err)
	}

	fmt.Println("Name       : ", jsn.Name)
	fmt.Println("Login      : ", jsn.Login)
	fmt.Println("ID         : ", jsn.ID)
	fmt.Println("Avatar URL : ", jsn.AvatarURL)
	fmt.Println("Profile URL: ", jsn.HTMLURL)
	fmt.Println("Company    : ", jsn.Company)
	fmt.Println("Blog       : ", jsn.Blog)
	fmt.Println("Bio        : ", jsn.Bio)
}
