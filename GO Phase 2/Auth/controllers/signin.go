package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var oauthGoogle *oauth2.Config
var RandomString = "random-string-" //random string can be generated using generater function

// func for initializing oauth
func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	oauthGoogle = &oauth2.Config{
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func Signin(w http.ResponseWriter, r *http.Request) {
	url := oauthGoogle.AuthCodeURL(RandomString)
	fmt.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
