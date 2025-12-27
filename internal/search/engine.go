package search

import "github.com/go-rod/rod"

type Profile struct {
	Name string
	URL  string
}

func Parse(page *rod.Page) ([]Profile, error) {
	links := page.MustElements(".profile")

	var profiles []Profile
	for _, link := range links {
		profiles = append(profiles, Profile{
			Name: link.MustText(),
			URL:  link.MustProperty("href").String(),
		})
	}

	return profiles, nil
}
