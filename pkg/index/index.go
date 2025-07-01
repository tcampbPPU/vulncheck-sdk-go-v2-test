package index

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/vulncheck-oss/sdk-go-v2/vulncheck"
)

func GetIndexInitialAccess() {
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
	resp, httpRes, err := client.IndicesAPI.IndexExploitsGet(auth).Execute()

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

func GetIndexWithCveFilter() {
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
	resp, httpRes, err := client.IndicesAPI.IndexInitialAccessGet(auth).Cve("CVE-2023-27350").Execute()

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

func GetIndexWithBotnetFilter() {
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
	resp, httpRes, err := client.IndicesAPI.IndexBotnetsGet(auth).Botnet("Fbot").Execute()

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

func GetIndexIpIntel() {
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
	resp, httpRes, err := client.IndicesAPI.IndexIpintel3dGet(auth).Country("Sweden").Id("c2").Execute()

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
