package main

import (
	"path/filepath"

	"humanflow-automation/internal/auth"
	"humanflow-automation/internal/browser"
	"humanflow-automation/internal/config"
	"humanflow-automation/internal/logger"
	"humanflow-automation/internal/search"
)

func main() {
	log := logger.New()
	cfg := config.Load("config.yaml")

	manager, err := browser.New(cfg.Browser.Headless)
	if err != nil {
		panic(err)
	}

	// 🔥 IMPORTANT: use absolute paths
	base, err := filepath.Abs(filepath.Join("mock", "html"))
	if err != nil {
		panic(err)
	}

	loginPath := "file:///" + filepath.Join(base, "login.html")
	searchPath := "file:///" + filepath.Join(base, "search.html")

	log.Info().Msg("starting mock automation")

	// Login flow
	if err := auth.Login(manager.Page, loginPath); err != nil {
		panic(err)
	}

	// Navigate to search page
	manager.Page.MustNavigate(searchPath)

	// Parse results
	profiles, _ := search.Parse(manager.Page)
	for _, p := range profiles {
		log.Info().
			Str("name", p.Name).
			Str("url", p.URL).
			Msg("profile discovered")
	}
}
