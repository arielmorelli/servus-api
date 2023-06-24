package filehandler

import (
	"testing"

	models "github.com/arielmorelli/servus-api/models"
	route "github.com/arielmorelli/servus-api/route"
)

// LoadRoutesFromFile
func TestLoadRoutesFromFile(t *testing.T) {
	// given
	route.Routes = make(models.Route)

	LoadRoutesFromFile("file_example.json")

	if len(route.Routes["a"]) != 2 {
		t.Fatalf("Route /a not added")
	}

	if len(route.Routes["b"]) != 1 {
		t.Fatalf("Route /a not added")
	}
}
