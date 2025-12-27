package auth

import "github.com/go-rod/rod"

func Login(page *rod.Page, loginPath string) error {
	// Navigate returns only error
	if err := page.Navigate(loginPath); err != nil {
		return err
	}

	page.MustElement("#email").MustInput("test@example.com")
	page.MustElement("#password").MustInput("password")
	page.MustElement("#loginBtn").MustClick()

	return nil
}
