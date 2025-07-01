package backup

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/vulncheck-oss/sdk-go-v2/vulncheck"
)

func GetIndexBackup() {
	configuration := vulncheck.NewConfiguration()
	configuration.Scheme = "https"
	configuration.Host = "api.vulncheck.com"

	client := vulncheck.NewAPIClient(configuration)

	token := os.Getenv("VULNCHECK_API_TOKEN")
	auth := context.WithValue(
		context.Background(),
		vulncheck.ContextAPIKeys,
		map[string]vulncheck.APIKey{
			"Bearer": {Key: token},
		},
	)
	resp, httpRes, err := client.EndpointsAPI.BackupIndexGet(auth, "mitre-cvelist-v5").Execute()

	if err != nil || httpRes.StatusCode != 200 {
		log.Fatal(err)
	}

	prettyJSON, err := json.MarshalIndent(resp.Data, "", "  ")
	if err != nil {
		log.Fatalf("Failed to generate JSON: %v", err)
		return
	}

	fmt.Println(string(prettyJSON))
}
