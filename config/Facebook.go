package config

import (
	"os"

	"golang.org/x/oauth2"
	facebookOAuth "golang.org/x/oauth2/facebook"
)

func GetFacebookOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
		ClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("FACEBOOK_REDIRECT_URL"),
		Endpoint:     facebookOAuth.Endpoint,
		Scopes:       []string{"public_profile", "email", "pages_show_list", "pages_read_engagement", "pages_manage_posts", "pages_messaging", "pages_manage_metadata"},
	}
}
