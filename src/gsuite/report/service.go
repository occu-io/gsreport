package report

import (
	"log"
	"net/http"

	admin "google.golang.org/api/admin/reports/v1"
)

func getService(client *http.Client) *admin.Service {
	srv, err := admin.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve directory client: %v", err)
	}

	return srv
}
