package footballdata_test

import (
	"fmt"
	"net/http"

	"github.com/icedream/go-footballdata"
)

func Example() {
	// Create client and tell it to use our API token
	client := new(footballdata.Client)

	// Tell it to use our API token (optional)
	client.SetToken("<insert your api token here>")

	// Get list of seasons...
	seasons, err := client.SoccerSeasons().Do()
	if err != nil {
		panic(err)
	}

	// ...and print them
	for _, season := range seasons {
		fmt.Println(season.Id, season.Caption)
	}
}

func ExampleClient() {
	// Create client
	client := footballdata.NewClient(http.DefaultClient)

	// Tell it to use our API token (optional)
	client.SetToken("<insert your api token here>")

	// Do something with the client instance...
	// Here we just fetch the listed soccer seasons on the API
	seasons, err := client.SoccerSeasons().Do()
	if err != nil {
		panic(err)
	}
	for _, season := range seasons {
		fmt.Println(season.Id, season.Caption)
	}
}

func ExampleClient_setTokenAndHttpClient() {
	// Create client
	client := new(footballdata.Client)

	// If you have an API token, you can tell the Client instance to use it
	client.SetToken("<insert your api token here>")

	// The Client instance also allows you to use your own HTTP client configuration
	client.SetHttpClient(&http.Client{
		Transport: &http.Transport{
			DisableCompression: true,
		}})
}
