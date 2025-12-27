package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type Manager struct {
	Browser *rod.Browser
	Page    *rod.Page
}

func New(headless bool) (*Manager, error) {
	u := launcher.New().
		Leakless(false).
		Headless(headless).
		MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()
	page := browser.MustPage()

	return &Manager{
		Browser: browser,
		Page:    page,
	}, nil
}
