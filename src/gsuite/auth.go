package gsuite

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	admin "google.golang.org/api/admin/directory/v1"
	report "google.golang.org/api/admin/reports/v1"
)

type Auth struct {
	credentialsFile string
	tokenFile       string
}

// Retrieve a token, saves the token, then returns the generated client.
func (au *Auth) getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tok, err := au.tokenFromFile()
	if err != nil {
		tok = au.getTokenFromWeb(config)
		au.saveToken(tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func (au *Auth) getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func (au *Auth) tokenFromFile() (*oauth2.Token, error) {
	f, err := os.Open(au.tokenFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func (au *Auth) saveToken(token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", au.tokenFile)
	f, err := os.OpenFile(au.tokenFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func GetClient(credentialsFile string, tokenFile string) *http.Client {
	var au Auth
	au.credentialsFile = credentialsFile
	au.tokenFile = tokenFile

	b, err := ioutil.ReadFile(au.credentialsFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, admin.AdminDirectoryUserScope, admin.AdminDirectoryGroupReadonlyScope, admin.AdminDirectoryUserSecurityScope, report.AdminReportsAuditReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := au.getClient(config)

	return client
}
