package config

import (
	"log"
	"os"
    "github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	//"golang.org/x/oauth2/google"
)

const Secret="secret"
type Config struct{
	GoogleLoginConfig oauth2.Config
}

var AppConfig Config

func GoogleConfig() oauth2.Config {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Some error occured. Err: %s", err)
    }

    AppConfig.GoogleLoginConfig = oauth2.Config{
        RedirectURL:  "http://localhost:8080/google_callback",
        ClientID:     Load("GOOGLE_CLIENT_ID"),
        ClientSecret: Load("GOOGLE_CLIENT_SECRET"),
        Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
            "https://www.googleapis.com/auth/userinfo.profile"},
        Endpoint: google.Endpoint,
    }

    return AppConfig.GoogleLoginConfig
}

//To load .env variables
func Load(key string) string{
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Some error occured. Err: %s", err)
    }
    return os.Getenv(key)
}