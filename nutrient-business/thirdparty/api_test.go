package thirdparty

import (
	"context"
	"os"
	"testing"

	"github.com/suhdev/nutrient/proton"
)

func TestEdmamSearchAPI(t *testing.T) {
	// This test only checks that the API server is up and that the response is
	// parsable based on the structs we have. However, more elaborate tests
	// must be included to ensure the completeness of test coverage
	t.Parallel()
	var (
		id, key string
		ok      bool
	)
	if id, ok = os.LookupEnv(proton.EnvThirdPartyAppID); !ok {
		t.Fatal("Could not find edmam app id")
	}
	if key, ok = os.LookupEnv(proton.EnvThirdPartyAppKey); !ok {
		t.Fatal("Could not find edmam app key")
	}

	ed := NewEdmamAPI(id, key)
	res, err := ed.Search(context.Background(), &proton.QueryRequest{
		Q: "chicken",
	})
	if err != nil {
		t.Fatalf("Did not expect an error, but go %v", err)
	}
	if len(res.Entries) == 0 {
		t.Fatal("Expected API to return non-zero results, but got none")
	}
}
