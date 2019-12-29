package directory

import (
	"log"
	"net/http"

	admin "google.golang.org/api/admin/directory/v1"
)

func getService(client *http.Client) *admin.Service {
	srv, err := admin.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve directory client: %v", err)
	}

	return srv
}

func getUserList(srv *admin.Service, query string) *admin.Users {
	r, err := srv.Users.List().Customer("my_customer").Query(query).MaxResults(500).OrderBy("email").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve users in domain: %v", err)
	}

	return r
}
